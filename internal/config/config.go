package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

type API struct {
	Address string `env:"API_ADDRESS,required" example:"localhost:8080"`
}

type Profiles struct {
	Address   string `env:"PROFILES_ADDRESS,required" example:"localhost:8081"`
	EnableTLS bool   `env:"PROFILES_ENABLE_TLS" envDefault:"false"`
}

type Auth struct {
	Address   string `env:"AUTH_ADDRESS,required" example:"localhost:8082"`
	EnableTLS bool   `env:"AUTH_ENABLE_TLS" envDefault:"false"`
}

type Config struct {
	API       API
	Profiles  Profiles
	Auth      Auth
	Namespace string `env:"NAMESPACE,required" envDefault:"gateway"`
}

func Load() (Config, error) {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		return Config{}, fmt.Errorf("parsing config: %w", err)
	}
	return cfg, nil
}
