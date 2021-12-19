package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes(app *fiber.App) {

	app.Use(cors.New())

	// api := app.Group("/api")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to my App test")
	})

	// v1 := api.Group("/v1"){
	// 	v1.Get("movies/:id",)
	// }
}
