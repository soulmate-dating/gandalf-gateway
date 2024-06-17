package http

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	echoPrometheus "github.com/globocom/echo-prometheus"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/soulmate-dating/gandalf-gateway/internal/app"
	"github.com/soulmate-dating/gandalf-gateway/internal/ports/http/middleware"
)

func NewServer(addr string, a app.ServiceLocator, namespace string) *http.Server {
	server := echo.New()
	s := &http.Server{Addr: addr, Handler: server}

	configMetrics := echoPrometheus.NewConfig()
	configMetrics.Namespace = namespace

	server.Use(echoPrometheus.MetricsMiddlewareWithConfig(configMetrics))
	server.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
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
