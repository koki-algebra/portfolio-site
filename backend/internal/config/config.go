package config

import (
	"context"

	"github.com/caarlos0/env/v6"
)

type EnvType string

const (
	Dev  EnvType = "dev"
	Prod EnvType = "prod"
)

var (
	Env = new(EnvConfig)
)

type EnvConfig struct {
	// Env
	EnvType EnvType `env:"ENV_TYPE" envDefault:"dev"`
	// Database
	DBHost     string `env:"DB_HOST"`
	DBPort     int    `env:"DB_PORT" envDefault:"5432"`
	DBDatabase string `env:"DB_DATABASE"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	// Server
	ServerPort         int    `env:"SERVER_PORT" envDefault:"80"`
	ServerAllowOrigins string `env:"SERVER_ALLOW_ORIGINS"`
	// Google Cloud
	GoogleClientID     string `env:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret string `env:"GOOGLE_CLIENT_SECRET"`
}

func Init(ctx context.Context) (err error) {
	if err = env.Parse(Env); err != nil {
		return
	}

	return
}
