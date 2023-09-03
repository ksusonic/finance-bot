package controller

import tele "gopkg.in/telebot.v3"

type Handler struct {
	Endpoint    string
	HandlerFunc tele.HandlerFunc
	Description string
	Middleware  []tele.MiddlewareFunc
}
