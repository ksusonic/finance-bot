package telegram

import (
	"strconv"

	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

type Context struct {
	bot    *tele.Bot
	chat   *tele.Chat
	logger *zap.SugaredLogger
}

func (c *Context) Logger() *zap.SugaredLogger {
	return c.logger
}

func (c *Context) Bot() *tele.Bot {
	return c.bot
}

func (c *Context) Recipient() string {
	return strconv.Itoa(int(c.chat.ID))
}

func (c *Context) Notify(action tele.ChatAction) {
	if err := c.bot.Notify(c, action); err != nil {
		c.logger.Errorf("unable to notify action: %s", action)
	}
}
