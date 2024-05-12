package http

import (
	"context"
	"log"
	"os"

	"golang.org/x/sync/errgroup"

	"github.com/soulmate-dating/gandalf-gateway/internal/app"
	"github.com/soulmate-dating/gandalf-gateway/internal/config"
	"github.com/soulmate-dating/gandalf-gateway/internal/graceful"
)

func Run(ctx context.Context, cfg config.Config, serviceLocator app.ServiceLocator) {
	httpServer := NewServer(cfg.API.Address, serviceLocator)

	eg, ctx := errgroup.WithContext(ctx)
	sigQuit := make(chan os.Signal, 1)
	eg.Go(graceful.CaptureSignal(ctx, sigQuit))
	eg.Go(RunServer(ctx, httpServer))
	if err := eg.Wait(); err != nil {
		log.Printf("gracefully shutting down the servers: %s\n", err.Error())
	}
	log.Println("servers were successfully shutdown")
}
