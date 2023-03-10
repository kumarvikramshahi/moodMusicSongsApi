package main

import (
	"moodMusicSongsApi/configs"
	"moodMusicSongsApi/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	//run database
	configs.DatabaseConnection()
	app.Use(logger.New())

	//routes
	routes.UserRoute(app)
	routes.SongRoute(app)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3001")
}
