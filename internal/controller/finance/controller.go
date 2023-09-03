package finance

import (
	"github.com/jmoiron/sqlx"
	tele "gopkg.in/telebot.v3"
)

type Controller struct {
	bot *tele.Bot
	db  *sqlx.DB
}

func NewController(bot *tele.Bot, db *sqlx.DB) *Controller {
	return &Controller{
		bot: bot,
		db:  db,
	}
}
