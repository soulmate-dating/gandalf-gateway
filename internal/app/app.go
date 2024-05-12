package app

import (
	"context"
	"log"
	"os"

	"golang.org/x/sync/errgroup"

	"github.com/soulmate-dating/gandalf-gateway/internal/app/clients/auth"
	clientcfg "github.com/soulmate-dating/gandalf-gateway/internal/app/clients/config"
	"github.com/soulmate-dating/gandalf-gateway/internal/app/clients/profiles"
	"github.com/soulmate-dating/gandalf-gateway/internal/config"
	"github.com/soulmate-dating/gandalf-gateway/internal/graceful"
	"github.com/soulmate-dating/gandalf-gateway/internal/ports/http"
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
		Address: cfg.Profiles.Address,
		UseSSL:  cfg.Profiles.EnableTLS,
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

func Run(ctx context.Context, cfg config.Config, serviceLocator ServiceLocator) {
	httpServer := http.NewServer(cfg.API.Address, serviceLocator)

	eg, ctx := errgroup.WithContext(ctx)
	sigQuit := make(chan os.Signal, 1)
	eg.Go(graceful.CaptureSignal(ctx, sigQuit))
	eg.Go(http.RunServer(ctx, httpServer))
	if err := eg.Wait(); err != nil {
		log.Printf("gracefully shutting down the servers: %s\n", err.Error())
	}
	log.Println("servers were successfully shutdown")
}
