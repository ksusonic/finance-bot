package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/ksusonic/finance-bot/internal/config"
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

	srv, err := service.NewBotService(cfg, log)
	if err != nil {
		log.Panicw("Service init error", "error", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv.Serve(ctx)
}
