package auth

import (
	"github.com/ksusonic/finance-bot/internal/telegram"

	tele "gopkg.in/telebot.v3"
)

type FilterFunction func(sender *tele.User, chat *tele.Chat) bool

type Controller struct {
	next            telegram.Controller
	filtrateMessage FilterFunction
}

func NewController(filter FilterFunction, next telegram.Controller) *Controller {
	return &Controller{next: next, filtrateMessage: filter}
}

func (c Controller) OnText(ctx telegram.Context, msg *tele.Message) {
	if !c.filtrateMessage(msg.Sender, msg.Chat) {
		c.next.OnText(ctx, msg)
	}
}

func (c Controller) OnCallback(ctx telegram.Context, cb *tele.Callback) {
	if !c.filtrateMessage(cb.Sender, nil) {
		c.next.OnCallback(ctx, cb)
	}
}
