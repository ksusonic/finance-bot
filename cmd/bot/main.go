package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/ksusonic/finance-bot/internal/config"
	"github.com/ksusonic/finance-bot/internal/controller/finance"
	"github.com/ksusonic/finance-bot/internal/logger"
	"github.com/ksusonic/finance-bot/internal/service"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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

	log.Debugf("connecting do database: %s", cfg.DatabaseDsn)
	db, err := sqlx.Connect("postgres", cfg.DatabaseDsn)
	if err != nil {
		log.Panicf("Database init error: %v", err)
	}

	log.Debug("initializing bot service")
	srv, err := service.NewBotService(cfg, log)
	if err != nil {
		log.Panicf("Service init error: %v", err)
	}

	log.Debug("registering controllers")
	srv.RegisterController(
		finance.NewController(srv.Bot(), db),
		// add new here
	)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv.Serve(ctx)
}
