package config

import (
	"github.com/caarlos0/env/v7"
)

type Config struct {
	Env        string `env:"ENV" envDefault:"dev"`
	Port       int    `env:"PORT" envDefault:"8080"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	DBHost     string `env:"DB_HOST"`
	DBPort     int    `env:"DB_PORT"`
	DBName     string `env:"DB_DBNAME"`
	SearchPath string `env:"SEARCH_PATH"`
	SSLMode    string `env:"SSL_MODE"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
