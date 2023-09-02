package middleware

import (
	"fmt"

	"github.com/samber/lo"
	tele "gopkg.in/telebot.v3"
)

func AllowForUsers(allowedUsers []int64) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			if lo.Contains(allowedUsers, c.Sender().ID) {
				return next(c)
			}
			return fmt.Errorf("user %d with username %s is not allowed", c.Sender().ID, c.Sender().Username)
		}
	}
}
