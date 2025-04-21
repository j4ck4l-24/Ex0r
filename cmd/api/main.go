package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/j4ck4l-24/Ex0r/internal/router"
	"github.com/j4ck4l-24/Ex0r/pkg/config"
	db "github.com/j4ck4l-24/Ex0r/pkg/database"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	router.ApiRoutes(app)

	if err := db.InitDB(); err != nil {
		log.Fatalf("DB init failed: %v", err)
	}

	_, appConfig, _ := config.Load()

	fmt.Println(app.Listen(fmt.Sprintf(":%v", appConfig.Port)))
}
