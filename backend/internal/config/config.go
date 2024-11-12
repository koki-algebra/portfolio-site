package config

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"github.com/caarlos0/env/v6"
)

type EnvType string

const (
	Dev  EnvType = "dev"
	Prod EnvType = "prod"
)

var (
	Env         = new(EnvConfig)
	FirebaseApp = new(firebase.App)
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
}

func Init(ctx context.Context) (err error) {
	if err := env.Parse(Env); err != nil {
		return err
	}

	FirebaseApp, err = firebase.NewApp(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
