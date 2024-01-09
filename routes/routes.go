package routes

import (
	"projectxiel/handlers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/post/:slug", handlers.GetSinglePost)
	app.Get("/posts", handlers.GetAllPosts)
	app.Static("/", "./public")
}
