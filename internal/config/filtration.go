package config

import "gopkg.in/telebot.v3"

type UserAllowFunc func(sender *telebot.User, chat *telebot.Chat) bool
