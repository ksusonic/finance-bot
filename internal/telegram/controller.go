package telegram

import tele "gopkg.in/telebot.v3"

type Controller interface {
	OnText(ctx Context, msg *tele.Message) error
	OnCallback(ctx Context, cb *tele.Callback) error
}
