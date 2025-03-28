package router

import (
	"github.com/gofiber/fiber/v2"
)


func ApiRoutes(app *fiber.App) {
    api := app.Group("/api",)
    api.Get("/testing", greet)
}


func greet(c *fiber.Ctx) error{
    return c.SendString("testing 123 hello")
}
