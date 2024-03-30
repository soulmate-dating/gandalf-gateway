package middleware

import (
	"github.com/labstack/echo/v4"
	"log"
)

func RecoveryMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("panic: %v", r)
			}
		}()

		return next(c)
	}
}
