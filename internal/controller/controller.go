package controller

import (
	"regexp"
	"strings"

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
	controller.commands[regexp.MustCompile("/listcommands")] = controller.listCommands
	return controller
}

func (c *Controller) listCommands(ctx telegram.Context, msg *telebot.Message) {
	var lines []string
	for command := range c.commands {
		lines = append(lines, command.String())
	}
	_, _ = ctx.Bot().Send(msg.Chat, strings.Join(lines, "\n"))
}

func (c *Controller) AddCommand(exp *regexp.Regexp, handler func(ctx telegram.Context, msg *telebot.Message)) {
	c.commands[exp] = handler
}

func (c *Controller) AddCommandWithArgs(
	exp *regexp.Regexp,
	handler func(ctx telegram.Context, msg *telebot.Message, args []string),
) {
	c.AddCommand(exp, func(ctx telegram.Context, msg *telebot.Message) {
		handler(ctx, msg, exp.FindAllStringSubmatch(msg.Text, -1)[0][1:])
	})
}

func (c *Controller) AddCallback(exp *regexp.Regexp, handler func(ctx telegram.Context, cb *telebot.Callback)) {
	c.callbacks[exp] = handler
}

func (c *Controller) OnText(ctx telegram.Context, msg *telebot.Message) {
	ctx.Notify(telebot.Typing)
	// todo state
	c.next.OnText(ctx, msg)
}

func (c *Controller) OnCallback(ctx telegram.Context, cb *telebot.Callback) {
	// todo process
	c.next.OnCallback(ctx, cb)
}
