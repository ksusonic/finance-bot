package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(cfg *zap.Config) (_ *zap.SugaredLogger, err error) {
	var logger *zap.Logger
	if cfg == nil {
		cfg.EncoderConfig = zap.NewProductionEncoderConfig()
		cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		logger, err = cfg.Build()
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		return
	}
	return logger.Sugar(), err
}
