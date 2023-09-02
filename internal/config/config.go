package config

import (
	"flag"
	"os"

	"go.uber.org/zap"
)

type Config struct {
	Token       string `yaml:"token"`
	DatabaseDsn string `yaml:"database_dsn"`

	ZapLoggerConfig *zap.Config `yaml:"logger"`
}

func NewConfig() (*Config, error) {
	cfg := Config{}
	defer flag.Parse()

	flag.StringVar(&cfg.DatabaseDsn, "database-dsn", "", "Database DSN")
	if token := os.Getenv("TOKEN"); token != "" {
		cfg.Token = token
	} else {
		panic("telegram \"TOKEN\" environment variable is not set")
	}
	// TODO
	return &cfg, nil
}
