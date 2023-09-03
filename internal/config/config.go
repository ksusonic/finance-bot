package config

import (
	"flag"
	"os"

	"github.com/caarlos0/env/v9"
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Token                string      `env:"TOKEN,required"`
	DatabaseDsn          string      `yaml:"database_dsn" env:"DATABASE_DSN"`
	ZapLoggerConfig      *zap.Config `yaml:"logger"`
	FiltrationConfigPath string      `yaml:"filtration_config_path"`

	FiltrationConfig FiltrationConfig `yaml:"-"`
	Telegram         *tele.Settings   `yaml:"-"`
}

type FiltrationConfig struct {
	AllowedUsers []int64 `yaml:"allowed_users"`
}

func NewConfig() (*Config, error) {
	cfg := Config{}

	configFile := flag.String("config", "config/dev.yaml", "Path to config")
	flag.Parse()

	// Open config file and unmarshall it to cfg
	file, err := os.Open(*configFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}

	if cfg.FiltrationConfigPath != "" {
		filtrationFile, err := os.Open(cfg.FiltrationConfigPath)
		if err != nil {
			return nil, err
		}
		defer filtrationFile.Close()

		decoder = yaml.NewDecoder(filtrationFile)
		err = decoder.Decode(&cfg.FiltrationConfig)
		if err != nil {
			return nil, err
		}
	}

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	cfg.Telegram = telegramConfig(cfg.Token)

	return &cfg, nil
}
