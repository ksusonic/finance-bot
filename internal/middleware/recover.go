package middleware

import (
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
	teleMiddleware "gopkg.in/telebot.v3/middleware"
)

func Recover(logger *zap.SugaredLogger) tele.MiddlewareFunc {
	return teleMiddleware.Recover(func(err error) {
		logger.Errorf("unexpected panic: %v", err)
	})
}
