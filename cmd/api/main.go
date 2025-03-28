package main

import (
	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/j4ck4l-24/Ex0r/internal/router"
)

func main() {
    app := fiber.New()

    router.ApiRoutes(app)

    // app.Use(cors.New())

    app.Listen(":8000")
}
