package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/soulmate-dating/gandalf-gateway/internal/ports/http/response"
	"log"
	"net/http"
)

func RecoveryMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("panic: %+v", r)
				err = c.JSON(http.StatusInternalServerError, response.Error("internal server error"))
			}
		}()

		return next(c)
	}
}
