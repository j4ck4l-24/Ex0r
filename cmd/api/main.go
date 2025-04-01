package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/j4ck4l-24/Ex0r/internal/router"
	"github.com/j4ck4l-24/Ex0r/pkg/config"
	"github.com/j4ck4l-24/Ex0r/pkg/database"
)

func main() {
	app := fiber.New()
	router.ApiRoutes(app)
	app.Use(cors.New())

	db.InitDB()

	_, appConfig, _ := config.Load()

	log.Fatalln(app.Listen(fmt.Sprintf(":%v", appConfig.Port)))
}
