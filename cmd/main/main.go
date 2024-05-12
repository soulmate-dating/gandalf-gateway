package main

import (
	"context"
	_ "github.com/soulmate-dating/gandalf-gateway/docs"
	"github.com/soulmate-dating/gandalf-gateway/internal/app"
	"github.com/soulmate-dating/gandalf-gateway/internal/config"
	"log"
)

const (
	httpPort = ":3000"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	svc := app.New(ctx, cfg)
	app.Run(ctx, cfg, svc)
}
