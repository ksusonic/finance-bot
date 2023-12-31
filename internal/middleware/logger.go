package middleware

import (
	"encoding/json"

	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

func Logger(logger *zap.SugaredLogger) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			data, _ := json.Marshal(c.Update())
			logger.Infof("got incoming event: %s", data)
			return next(c)
		}
	}
}
