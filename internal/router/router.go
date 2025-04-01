package router

import (
	// "fmt"
	"github.com/gofiber/fiber/v2"
    "github.com/j4ck4l-24/Ex0r/internal/handlers/auth"
)


func ApiRoutes(app *fiber.App) {
    api := app.Group("/api",)
    app.Get("/healthcheck", healthcheck)

    auth_api := api.Group("/auth")
    auth_api.Post("/login", auth.LoginAttempt)
    auth_api.Post("/register",healthcheck)
    auth_api.Post("/logout",healthcheck)

    // I don't think so there is need of admin api as of now as we can implement token based role system
    // admin := api.Group("/admin")
    // admin.Get("/",healthcheck)

}


func healthcheck(c *fiber.Ctx) error{
    return c.SendString("Server is Running")
}
