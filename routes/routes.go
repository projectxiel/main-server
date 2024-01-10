package routes

import (
	"projectxiel/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)


func Routes(app *fiber.App) {
	app.Get("/post/:slug", handlers.GetSinglePost)
	app.Get("/posts", handlers.GetAllPosts)
	app.Static("/", "./public")
	app.Get("/swagger/*", swagger.HandlerDefault)
}
