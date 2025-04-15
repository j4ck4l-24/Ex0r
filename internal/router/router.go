package router

import (
	// "fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/j4ck4l-24/Ex0r/internal/handlers/auth"
	"github.com/j4ck4l-24/Ex0r/internal/handlers/challenges"
	"github.com/j4ck4l-24/Ex0r/internal/handlers/flags"
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
	api.Get("/challenges/:id", middleware.Protected(), challenges.GetChallenge)
	api.Post("/challenges", middleware.Protected(), middleware.AdminOnly(), challenges.CreateChallenge)
	api.Patch("/challenges/:id", middleware.Protected(), middleware.AdminOnly(), challenges.UpdateChallenge)
	api.Delete("/challenges/:id", middleware.Protected(), middleware.AdminOnly(), challenges.DeleteChallenge)

	api.Get("/flags", middleware.Protected(), middleware.AdminOnly(), flags.GetAllFlag)
	api.Get("/flags/:id", middleware.Protected(), middleware.AdminOnly(), flags.GetFlag)
	api.Post("/flags", middleware.Protected(), middleware.AdminOnly(), flags.CreateFlag)
	api.Patch("/flags/:id", middleware.Protected(), middleware.AdminOnly(), flags.UpdateFlag)
	api.Delete("/flags/:id", middleware.Protected(), middleware.AdminOnly(), flags.DeleteFlag)
}

func healthcheck(c *fiber.Ctx) error {
	return c.SendString("Server is Running")
}
