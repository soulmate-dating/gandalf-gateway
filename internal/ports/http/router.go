package http

import (
	"github.com/labstack/echo/v4"
	"github.com/soulmate-dating/gandalf-gateway/internal/app"
	"github.com/soulmate-dating/gandalf-gateway/internal/ports/http/api/auth"
	"github.com/soulmate-dating/gandalf-gateway/internal/ports/http/api/profiles"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// RegisterRoutes registers all routes
// @title Gandalf API Gateway
// @version 1.0
// @description Authentication and profiles gateway
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost
// @BasePath /api/v0
func RegisterRoutes(e *echo.Echo, a app.ServiceLocator) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	group := e.Group("/api/v0")

	auth.RegisterRoutes(group, a)
	profiles.RegisterRoutes(group, a)
}
