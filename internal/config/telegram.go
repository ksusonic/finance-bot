package config

import (
	"time"

	tele "gopkg.in/telebot.v3"
)

func telegramConfig(token string) *tele.Settings {
	return &tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
		// TODO: load telegram config
	}
}
