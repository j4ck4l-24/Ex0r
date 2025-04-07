package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/j4ck4l-24/Ex0r/internal/router"
	"github.com/j4ck4l-24/Ex0r/pkg/config"
)

func main() {
	app := fiber.New()
	router.ApiRoutes(app)
	app.Use(cors.New())

	_, appConfig, _ := config.Load()

	fmt.Println(app.Listen(fmt.Sprintf(":%v", appConfig.Port)))
}
