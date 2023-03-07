package main

import (
	"moodMusicSongsApi/configs"
	"moodMusicSongsApi/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//run database
	configs.ConnectDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	//routes
	routes.UserRoute(app)

	app.Listen(":3001")
}
