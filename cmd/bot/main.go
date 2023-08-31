package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ksusonic/finance-bot/internal/config"
	"github.com/ksusonic/finance-bot/internal/telegram"
)

func main() {
	cfg := config.NewConfig()

	b := telegram.NewBot(cfg.Token)

	go b.Start()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown ...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	botStop := make(chan struct{}, 1)
	go func() {
		b.Stop()
		botStop <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		log.Println(ctx.Err())
	case <-botStop:
		log.Println("Bot stopped gracefully")
	}
}
