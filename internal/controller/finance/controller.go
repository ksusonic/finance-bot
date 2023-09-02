package finance

import (
	"github.com/ksusonic/finance-bot/internal/telegram"
	tele "gopkg.in/telebot.v3"
)

type Controller struct {
}

type TransactionStorage interface {
}

func NewController(txStorage TransactionStorage) *Controller {
	return &Controller{}
}

func (c Controller) OnText(ctx telegram.Context, msg *tele.Message) error {
	//TODO implement me
	panic("implement me")
	return nil
}

func (c Controller) OnCallback(ctx telegram.Context, cb *tele.Callback) error {
	//TODO implement me
	panic("implement me")
	return nil
}
