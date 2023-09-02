package command

import (
	"regexp"

	"github.com/ksusonic/finance-bot/internal/telegram"

	"gopkg.in/telebot.v3"
)

type Controller struct {
	commands  map[*regexp.Regexp]func(ctx telegram.Context, msg *telebot.Message)
	callbacks map[*regexp.Regexp]func(ctx telegram.Context, msg *telebot.Callback)
	next      telegram.Controller
}

func NewController(next telegram.Controller) *Controller {
	controller := &Controller{
		commands:  make(map[*regexp.Regexp]func(ctx telegram.Context, msg *telebot.Message)),
		callbacks: make(map[*regexp.Regexp]func(ctx telegram.Context, msg *telebot.Callback)),
		next:      next,
	}
	return controller
}

func (c *Controller) OnText(ctx telegram.Context, msg *telebot.Message) error {
	ctx.Notify(telebot.Typing)
	// todo state
	c.next.OnText(ctx, msg)
	return nil
}

func (c *Controller) OnCallback(ctx telegram.Context, cb *telebot.Callback) error {
	// todo process
	c.next.OnCallback(ctx, cb)
	return nil
}
