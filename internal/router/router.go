package router

import (
	// "fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/j4ck4l-24/Ex0r/internal/handlers/auth"
	"github.com/j4ck4l-24/Ex0r/internal/handlers/challenges"
	"github.com/j4ck4l-24/Ex0r/internal/middleware"
)

func ApiRoutes(app *fiber.App) {
	api := app.Group("/api")
	app.Get("/healthcheck", healthcheck)

	auth_api := api.Group("/auth")
	auth_api.Post("/login", auth.LoginAttempt)
	auth_api.Post("/register", auth.RegisterAttempt)
	auth_api.Post("/logout", healthcheck)

	// I don't think so there is need of admin api as of now as we can implement token based role system
	// admin := api.Group("/admin")
	// admin.Get("/",healthcheck)

	api.Get("/challenges", middleware.Protected(), challenges.GetAllChallenges)
	api.Get("/challenge/:id", middleware.Protected(), challenges.GetChallenge)
	api.Post("/challenge", middleware.Protected(), middleware.AdminOnly(), challenges.CreateChallenge)
	api.Patch("/challenge/:id", middleware.Protected(), middleware.AdminOnly(), challenges.UpdateChallenge)
    api.Delete("/challenge/:id", middleware.Protected(), middleware.AdminOnly(), challenges.DeleteChallenge)
}

func healthcheck(c *fiber.Ctx) error {
	return c.SendString("Server is Running")
}
