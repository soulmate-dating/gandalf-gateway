package main

import (
	"context"
	"github.com/soulmate-dating/gandalf-gateway/internal/app"
	"github.com/soulmate-dating/gandalf-gateway/internal/graceful"
	"github.com/soulmate-dating/gandalf-gateway/internal/ports/http"
	"golang.org/x/sync/errgroup"
	"os"

	"log"
)

const (
	httpPort = ":3000"
)

func main() {
	ctx := context.Background()

	serviceLocator := app.NewServiceLocator()

	httpServer := http.NewServer(httpPort, serviceLocator)

	eg, ctx := errgroup.WithContext(ctx)

	sigQuit := make(chan os.Signal, 1)
	eg.Go(graceful.CaptureSignal(ctx, sigQuit))
	// run http server
	eg.Go(http.RunServer(ctx, httpServer))

	if err := eg.Wait(); err != nil {
		log.Printf("gracefully shutting down the servers: %s\n", err.Error())
	}
	log.Println("servers were successfully shutdown")
}
