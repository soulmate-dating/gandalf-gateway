package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v6"
)

type Postgres struct {
	Host              string        `env:"POSTGRES_HOST" example:"localhost"`
	Port              int           `env:"POSTGRES_PORT" envDefault:"5432"`
	User              string        `env:"POSTGRES_USER" example:"glimpse"`
	Password          string        `env:"POSTGRES_PASSWORD" example:"password"`
	Database          string        `env:"POSTGRES_DB" example:"glimpse"`
	SSLMode           string        `env:"POSTGRES_SSL_MODE" envDefault:"disable"`
	ConnectionTimeout time.Duration `env:"POSTGRES_CONNECTION_TIMEOUT" envDefault:"60s"`
}

type API struct {
	Address string `env:"API_ADDRESS" envDefault:"localhost:8080"`
}

type Profiles struct {
	Address   string `env:"PROFILES_ADDRESS" envDefault:"localhost:8081"`
	EnableTLS bool   `env:"PROFILES_ENABLE_TLS" envDefault:"false"`
}

type Auth struct {
	Address   string `env:"AUTH_ADDRESS" envDefault:"localhost:8082"`
	EnableTLS bool   `env:"AUTH_ENABLE_TLS" envDefault:"false"`
}

type Config struct {
	Postgres Postgres
	API      API
	Profiles Profiles
	Auth     Auth
}

func Load() (Config, error) {
	var cfg Config
	err := env.Parse(cfg)
	if err != nil {
		return Config{}, fmt.Errorf("parsing config: %w", err)
	}
	return cfg, nil
}
