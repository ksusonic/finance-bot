package finance

import (
	"github.com/ksusonic/finance-bot/internal/telegram"
	tele "gopkg.in/telebot.v3"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (c Controller) OnText(ctx telegram.Context, msg *tele.Message) {
	//TODO implement me
	panic("implement me")
}

func (c Controller) OnCallback(ctx telegram.Context, cb *tele.Callback) {
	//TODO implement me
	panic("implement me")
}
