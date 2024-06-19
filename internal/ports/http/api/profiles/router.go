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
	r.GET("/profile/full", getFullProfile(locator.Profiles()))
	r.GET("/profile/recommendation", getRecommendation(locator.Profiles()))

	r.POST("/prompts/text", createPrompt(locator.Profiles()))
	r.PUT("/prompts/text/:prompt_id", updatePrompt(locator.Profiles()))
	r.DELETE("/prompts/text/:prompt_id", deletePrompt(locator.Profiles()))

	r.POST("/prompts/file", createFilePrompt(locator.Profiles()))
	r.PUT("/prompts/file/:prompt_id", updateFilePrompt(locator.Profiles()))
	r.DELETE("/prompts/file/:prompt_id", deletePrompt(locator.Profiles()))

	r.GET("/prompts", getPrompts(locator.Profiles()))
}
