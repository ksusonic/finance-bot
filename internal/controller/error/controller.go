package error

import (
	"github.com/ksusonic/finance-bot/internal/telegram"

	tele "gopkg.in/telebot.v3"
)

type Controller struct {
	next telegram.Controller
}

func NewController(next telegram.Controller) *Controller {
	return &Controller{next: next}
}

func (c *Controller) OnText(ctx telegram.Context, msg *tele.Message) {
	c.captureError(ctx, func() { c.next.OnText(ctx, msg) })
}

func (c *Controller) OnCallback(ctx telegram.Context, cb *tele.Callback) {
	c.captureError(ctx, func() { c.next.OnCallback(ctx, cb) })
}

func (c *Controller) captureError(ctx telegram.Context, f func()) {
	defer func() {
		if err := recover(); err != nil {
			ctx.Logger().Errorw(
				"caught controller panic",
				"panic", err,
			)
		}
	}()
	f()
}
