package logger

import (
	"github.com/ksusonic/finance-bot/internal/telegram"

	tele "gopkg.in/telebot.v3"
)

const (
	eventMsg   = "got event"
	typeKey    = "type"
	fromKey    = "from"
	payloadKey = "payload"
)

type Controller struct {
	next telegram.Controller
}

func NewController(next telegram.Controller) *Controller {
	return &Controller{next: next}
}

func (c Controller) OnText(ctx telegram.Context, msg *tele.Message) {
	ctx.Logger().Debugw(eventMsg,
		typeKey, "text",
		fromKey, msg.Sender.Username,
		payloadKey, msg.Text,
	)
	c.next.OnText(ctx, msg)
}

func (c Controller) OnCallback(ctx telegram.Context, cb *tele.Callback) {
	ctx.Logger().Debugw(eventMsg,
		typeKey, "callback",
		fromKey, cb.Sender.Username,
		payloadKey, cb.Data,
	)
	c.next.OnCallback(ctx, cb)
}
