package service

import (
	"context"
	"time"

	"github.com/ksusonic/finance-bot/internal/config"
	"github.com/ksusonic/finance-bot/internal/controller/command"
	"github.com/ksusonic/finance-bot/internal/middleware"
	"github.com/ksusonic/finance-bot/internal/telegram"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

type Service struct {
	logger     *zap.SugaredLogger
	bot        *tele.Bot
	controller telegram.Controller
	db         *sqlx.DB
}

func NewBotService(
	cfg *config.Config,
	logger *zap.SugaredLogger,
	db *sqlx.DB,
	logicController telegram.Controller,
) (*Service, error) {
	pref := tele.Settings{
		Token:  cfg.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	b, err := tele.NewBot(pref)
	if err != nil {
		return nil, err
	}
	logger.Infof("logged in telegram as: %s", b.Me.Username)

	b.Use(middleware.Recover(logger.Named("recover")))
	b.Use(middleware.Logger(logger.Named("access-log")))
	if len(cfg.FiltrationConfig.AllowedUsers) > 0 {
		logger.Debugf("Allowed users: %s")
		b.Use(middleware.AllowForUsers(cfg.FiltrationConfig.AllowedUsers))
	} else {
		logger.Warn("No users whitelist provided")
	}

	var ctrl = command.NewController(logicController)
	{
		b.Handle(tele.OnText, func(c tele.Context) error {
			return ctrl.OnText(telegram.Context{}, c.Message())
		})
	}

	return &Service{
		logger:     logger,
		bot:        b,
		controller: ctrl,
		db:         db,
	}, nil
}

func (s *Service) Serve(ctx context.Context) {
	go s.bot.Start()
	<-ctx.Done()
	s.bot.Stop()
}
