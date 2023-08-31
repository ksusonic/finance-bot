package telegram

import (
	"time"

	tele "gopkg.in/telebot.v3"
)

type Bot struct {
	*tele.Bot
}

func NewBot(token string) *Bot {
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		panic(err)
	}
	initHandles(b)
	return &Bot{initHandles(b)}
}

func initHandles(b *tele.Bot) *tele.Bot {
	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Hello, I'm finance bot!")
	})

	return b
}
