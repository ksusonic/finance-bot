package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/ksusonic/finance-bot/internal/config"
	"github.com/ksusonic/finance-bot/internal/telegram"
)

func main() {
	cfg := config.NewConfig()

	b, err := telegram.NewBot(cfg.Token)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sqlx.Connect("postgres", cfg.DatabaseDsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//var (
	//	transactionStorage = storage.NewTransactionsStorage(db)
	//	userStorage        = storage.NewUsersStorage(db)
	//	chatStorage        = storage.NewChatsStorage(db)
	//)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	go func(ctx context.Context) {
		log.Println("Bot started")
		b.Start()
		log.Println("Bot stopped")
	}(ctx)

	<-ctx.Done()
	log.Println("Stopping...")

	cancelCtx, cancel2 := context.WithTimeout(ctx, 5*time.Second)
	defer cancel2()

	botStop := make(chan struct{}, 1)
	go func() {
		b.Stop()
		botStop <- struct{}{}
	}()
	select {
	case <-cancelCtx.Done():
		log.Println("Bot graceful shutdown timed out")
	case <-botStop:
		log.Println("Bot graceful shutdown")
	}
}
