package app

import (
	"fmt"
	"github.com/soulmate-dating/gandalf-gateway/internal/app/clients/auth"
	"github.com/soulmate-dating/gandalf-gateway/internal/app/clients/profile"
	"log"
)

var (
	ErrForbidden = fmt.Errorf("forbidden")
)

type ServiceLocator interface {
	Profiles() profile.ProfileServiceClient
	Auth() auth.AuthServiceClient
}

type locator struct {
	profileServiceClient profile.ProfileServiceClient
	authServiceClient    auth.AuthServiceClient
}

func (l *locator) Profiles() profile.ProfileServiceClient {
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
	profileServiceClient, err := profile.NewServiceClient()
	if err != nil {
		log.Fatalf("could not connect to profile service: %s", err.Error())
	}
	return &locator{
		authServiceClient:    authServiceClient,
		profileServiceClient: profileServiceClient,
	}
}
