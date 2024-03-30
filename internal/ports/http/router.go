package http

import (
	"github.com/labstack/echo/v4"
	"github.com/soulmate-dating/gandalf-gateway/internal/app"
	"github.com/soulmate-dating/gandalf-gateway/internal/ports/http/api/auth"
	"github.com/soulmate-dating/gandalf-gateway/internal/ports/http/api/profiles"
)

func RegisterRoutes(e *echo.Echo, a app.ServiceLocator) {
	group := e.Group("/api/v0")

	auth.RegisterRoutes(group, a)
	profiles.RegisterRoutes(group, a)
}
