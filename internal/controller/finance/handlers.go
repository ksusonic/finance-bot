package finance

import (
	"github.com/ksusonic/finance-bot/internal/controller"

	tele "gopkg.in/telebot.v3"
)

func (c *Controller) Handlers() []controller.Handler {
	return []controller.Handler{
		{
			Endpoint: "/test",
			HandlerFunc: func(c tele.Context) error {
				return c.Send(
					"Hello! I can help you to manage your finances 😎\n" +
						"Just send me a command /help to see what I can do",
				)
			},
			Description: "Первый тестовый контроллер",
		},
	}
}
