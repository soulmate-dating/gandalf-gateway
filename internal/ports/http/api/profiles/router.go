package profiles

import (
	"github.com/labstack/echo/v4"
	"github.com/soulmate-dating/gandalf-gateway/internal/app"
	"github.com/soulmate-dating/gandalf-gateway/internal/ports/http/middleware"
)

func RegisterRoutes(g *echo.Group, locator app.ServiceLocator) {
	r := g.Group("/users/:user_id")
	r.Use(middleware.InitAuthMiddleWare(locator.Auth()))
	r.POST("/profile", createProfile(locator.Profiles()))
	r.PUT("/profile", updateProfile(locator.Profiles()))
	r.GET("/profile", getProfile(locator.Profiles()))
	r.GET("/recommendation", getRecommendation(locator.Profiles()))
}
