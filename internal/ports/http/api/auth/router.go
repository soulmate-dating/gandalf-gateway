package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/soulmate-dating/gandalf-gateway/internal/app"
)

func RegisterRoutes(g *echo.Group, locator app.ServiceLocator) {
	r := g.Group("/auth")
	authClient := locator.Auth()
	r.POST("/signup", signup(authClient))
	r.POST("/login", login(authClient))
	r.GET("/refresh", refresh(authClient))
	r.POST("/logout", logout(authClient))
}
