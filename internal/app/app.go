package app

import (
	"fmt"
	"github.com/soulmate-dating/gandalf-gateway/internal/app/clients/auth"
	"github.com/soulmate-dating/gandalf-gateway/internal/app/clients/profiles"
	"log"
)

var (
	ErrForbidden = fmt.Errorf("forbidden")
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

func NewServiceLocator() ServiceLocator {
	authServiceClient, err := auth.NewServiceClient()
	if err != nil {
		log.Fatalf("could not connect to auth service: %s", err.Error())
	}
	profileServiceClient, err := profiles.NewServiceClient()
	if err != nil {
		log.Fatalf("could not connect to profiles service: %s", err.Error())
	}
	return &locator{
		authServiceClient:    authServiceClient,
		profileServiceClient: profileServiceClient,
	}
}
