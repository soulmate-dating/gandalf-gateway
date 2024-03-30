package middleware

import (
	"github.com/labstack/echo/v4"
	"log"
	"time"
)

func LoggerMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()

		log.Printf("-- received request -- | protocol: HTTP | method: %s | path: %s\n", c.Request().Method, c.Request().URL.Path)
		defer func() {
			latency := time.Since(start)
			log.Printf("-- handled request -- | protocol: HTTP | status: %d | latency: %+v | method: %s | path: %s\n", c.Response().Status, latency, c.Request().Method, c.Request().URL.Path)
		}()
		return next(c)
	}
}
