package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/pokervarino27/go-base/api/routes"
)

func main() {
	var port string

	app := fiber.New(fiber.Config{
		ReadBufferSize: 9000,
	})

	//Middlewares
	app.Use(recover.New())
	app.Use(logger.New())

	//Services
	routes.SetupRoutes(app)
	if port = os.Getenv("HTTP_PORT"); port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen(":" + port))
}
