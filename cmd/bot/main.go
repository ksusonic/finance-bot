package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/jmoiron/sqlx"
	"github.com/ksusonic/finance-bot/internal/config"
	"github.com/ksusonic/finance-bot/internal/controller/finance"
	"github.com/ksusonic/finance-bot/internal/logger"
	"github.com/ksusonic/finance-bot/internal/service"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic("Config error: " + err.Error())
	}
	log, err := logger.NewLogger(cfg.ZapLoggerConfig)
	if err != nil {
		panic("Logger error: " + err.Error())
	}

	db, err := sqlx.Connect("postgres", cfg.DatabaseDsn)
	if err != nil {
		log.Panicf("Database init error: %v", err)
	}

	financeController := finance.NewController(db)
	srv, err := service.NewBotService(
		cfg,
		log,
		db,
		financeController,
	)
	if err != nil {
		log.Panicf("Service init error: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv.Serve(ctx)
}
