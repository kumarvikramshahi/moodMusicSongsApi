package routes

import (
	"moodMusicSongsApi/controllers"

	"github.com/gofiber/fiber/v2"
)

func SongRoute(app *fiber.App) {
	app.Get("/song/:emotion/", controllers.GetManySongs)
}
