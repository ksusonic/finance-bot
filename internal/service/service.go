package service

import (
	"context"

	"github.com/ksusonic/finance-bot/internal/config"
	"github.com/ksusonic/finance-bot/internal/controller"
	"github.com/ksusonic/finance-bot/internal/middleware"

	"github.com/samber/lo"
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

type Service struct {
	bot       *tele.Bot
	logger    *zap.SugaredLogger
	endpoints map[string]string // endpoint to description
}

func NewBotService(
	cfg *config.Config,
	logger *zap.SugaredLogger,
) (*Service, error) {
	b, err := tele.NewBot(*cfg.Telegram)
	if err != nil {
		return nil, err
	}
	logger.Infof("logged in telegram as: %s", b.Me.Username)

	b.Use(middleware.Recover(logger.Named("recover")))
	b.Use(middleware.Logger(logger.Named("access-log")))

	if len(cfg.FiltrationConfig.AllowedUsers) > 0 {
		logger.Debugf("Allowed users: %d", len(cfg.FiltrationConfig.AllowedUsers))
		b.Use(middleware.AllowForUsers(cfg.FiltrationConfig.AllowedUsers))
	} else {
		logger.Warn("No users whitelist provided")
	}

	return &Service{
		logger:    logger,
		bot:       b,
		endpoints: make(map[string]string),
	}, nil
}

func (s *Service) Serve(ctx context.Context) {
	if len(s.endpoints) == 0 {
		s.logger.Panic("No controllers registered")
	}

	allCommands, err := s.initCommands()
	if err != nil {
		s.logger.Errorf("Failed to init commands: %v", err)
	}
	s.initSysHandlers(allCommands)

	go s.bot.Start()
	s.logger.Info("Started bot")
	<-ctx.Done()
	s.bot.Stop()
}

func (s *Service) Bot() *tele.Bot {
	return s.bot
}

func (s *Service) RegisterController(controllers ...Controller) {
	lo.ForEach(controllers, func(c Controller, _ int) {
		lo.ForEach(c.Handlers(), func(h controller.Handler, _ int) {
			if _, ok := s.endpoints[h.Endpoint]; ok {
				s.logger.Panicf("Endpoint %s registered twice", h.Endpoint)
			}
			s.endpoints[h.Endpoint] = h.Description
			s.bot.Handle(h.Endpoint, h.HandlerFunc, h.Middleware...)
		})
	})
}

type Controller interface {
	Handlers() []controller.Handler
}
