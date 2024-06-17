package app

import (
	"context"
	"github.com/soulmate-dating/gandalf-gateway/internal/app/clients/auth"
	clientcfg "github.com/soulmate-dating/gandalf-gateway/internal/app/clients/common"
	"github.com/soulmate-dating/gandalf-gateway/internal/app/clients/profiles"
	"github.com/soulmate-dating/gandalf-gateway/internal/config"
	"log"
)

type ServiceLocator interface {
	Profiles() profiles.ProfileServiceClient
	Auth() auth.AuthServiceClient
}

type locator struct {
	profileServiceClient profiles.ProfileServiceClient
	authServiceClient    auth.AuthServiceClient
}

func (l *locator) Profiles() profiles.ProfileServiceClient {
	return l.profileServiceClient
}

func (l *locator) Auth() auth.AuthServiceClient {
	return l.authServiceClient
}

func New(_ context.Context, cfg config.Config) ServiceLocator {
	authServiceClient, err := auth.NewServiceClient(clientcfg.Config{
		Address: cfg.Auth.Address,
		UseSSL:  cfg.Auth.EnableTLS,
	})
	if err != nil {
		log.Fatalf("could not connect to auth service: %s", err.Error())
	}
	profileServiceClient, err := profiles.NewServiceClient(clientcfg.Config{
		Address: cfg.Profiles.Address,
		UseSSL:  cfg.Profiles.EnableTLS,
	})
	if err != nil {
		log.Fatalf("could not connect to profiles service: %s", err.Error())
	}
	return &locator{
		authServiceClient:    authServiceClient,
		profileServiceClient: profileServiceClient,
	}
}
