package middleware

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/soulmate-dating/gandalf-gateway/internal/app"
	"github.com/soulmate-dating/gandalf-gateway/internal/app/clients/auth"
	"github.com/soulmate-dating/gandalf-gateway/internal/ports/http/response"
	"net/http"
)

func InitAuthMiddleWare(client auth.AuthServiceClient) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			accessToken := c.Request().Header.Get("Access-Token")
			if accessToken == "" {
				return c.JSON(http.StatusForbidden, response.Error(errors.New("access token not provided")))
			}

			_, err := client.Validate(
				c.Request().Context(),
				&auth.ValidateRequest{AccessToken: accessToken},
			)
			if err != nil {
				switch {
				case errors.Is(err, app.ErrForbidden):
					return c.JSON(http.StatusForbidden, response.Error(err))
				default:
					return c.JSON(http.StatusInternalServerError, response.Error(err))
				}
			}
			return next(c)
		}
	}
}