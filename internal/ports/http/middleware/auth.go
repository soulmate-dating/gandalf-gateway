package middleware

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/soulmate-dating/gandalf-gateway/internal/app/clients/auth"
	"github.com/soulmate-dating/gandalf-gateway/internal/ports/http/response"
	"net/http"
	"strings"
)

var (
	ErrWrongAuthHeaderFormat = errors.New("wrong authorization header format")
	ErrorMissingAccessToken  = errors.New("access token not provided")
)

const (
	BearerPrefix = "Bearer "
	AuthIDKey    = "auth_id"
)

func InitAuthMiddleWare(client auth.AuthServiceClient) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if !strings.HasPrefix(authHeader, BearerPrefix) {
				return c.JSON(
					http.StatusForbidden,
					response.Error(ErrWrongAuthHeaderFormat.Error()),
				)
			}
			accessToken := authHeader[len(BearerPrefix):]
			if accessToken == "" {
				return c.JSON(
					http.StatusForbidden,
					response.Error(ErrorMissingAccessToken.Error()),
				)
			}
			res, err := client.Validate(
				c.Request().Context(),
				&auth.ValidateRequest{AccessToken: accessToken},
			)
			if err != nil {
				return response.MapError(c, err)
			}
			c.Set(AuthIDKey, res.GetId())
			return next(c)
		}
	}
}
