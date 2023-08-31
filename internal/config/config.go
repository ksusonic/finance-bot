package config

import (
	"flag"
	"os"
)

type Config struct {
	Address     string `yaml:"address"`
	Token       string `yaml:"token"`
	DatabaseDsn string `yaml:"database_dsn"`
}

func NewConfig() *Config {
	cfg := Config{}
	defer flag.Parse()
	flag.StringVar(&cfg.Address, "address", ":8080", "Address to listen on")
	flag.StringVar(&cfg.DatabaseDsn, "database-dsn", "", "Database DSN")
	if token := os.Getenv("TOKEN"); token != "" {
		cfg.Token = token
	} else {
		panic("telegram \"TOKEN\" environment variable is not set")
	}

	return &cfg
}
