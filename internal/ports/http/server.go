package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/soulmate-dating/gandalf-gateway/internal/app"
	"github.com/soulmate-dating/gandalf-gateway/internal/ports/http/middleware"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func NewServer(port string, a app.ServiceLocator) *http.Server {
	server := echo.New()
	s := &http.Server{Addr: port, Handler: server}

	server.Use(middleware.LoggerMiddleWare)
	server.Use(middleware.RecoveryMiddleWare)

	RegisterRoutes(server, a)

	return s
}

func RunServer(ctx context.Context, server *http.Server) func() error {
	return func() error {
		log.Printf("starting http server, listening on %s\n", server.Addr)
		defer log.Printf("close http server listening on %s\n", server.Addr)

		errCh := make(chan error)

		defer func() {
			shCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
			defer cancel()

			if err := server.Shutdown(shCtx); err != nil {
				log.Printf("can't close http server listening on %s: %s", server.Addr, err.Error())
			}

			close(errCh)
		}()

		go func() {
			if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
				errCh <- err
			}
		}()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-errCh:
			return fmt.Errorf("http server can't listen and serve requests: %w", err)
		}
	}
}
