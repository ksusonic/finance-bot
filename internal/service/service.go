package service

import (
	"context"
	"time"

	"github.com/ksusonic/finance-bot/internal/config"
	"github.com/ksusonic/finance-bot/internal/controller"
	"github.com/ksusonic/finance-bot/internal/controller/finance"
	"github.com/ksusonic/finance-bot/internal/telegram"

	authController "github.com/ksusonic/finance-bot/internal/controller/auth"
	errorController "github.com/ksusonic/finance-bot/internal/controller/error"
	loggerController "github.com/ksusonic/finance-bot/internal/controller/logger"

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
) (*Service, error) {
	// telegram
	pref := tele.Settings{
		Token:  cfg.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	b, err := tele.NewBot(pref)
	if err != nil {
		return nil, err
	}
	logger.Infof("logged in telegram as: %s", b.Me.Username)

	// database
	logger.Debugf("connecting to postgres storage: %s", cfg.DatabaseDsn)
	db, err := sqlx.Connect("postgres", cfg.DatabaseDsn)
	if err != nil {
		return nil, err
	}

	var ctrl = loggerController.NewController(
		authController.NewController(
			func(sender *tele.User, chat *tele.Chat) bool {
				return true // TODO
			},
			errorController.NewController(
				controller.NewController(
					finance.NewController(),
				),
			),
		),
	)

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
