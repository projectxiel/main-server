package routes

import (
	"projectxiel/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func Routes(app *fiber.App) {
	app.Static("/", "./public")
	app.Static("/static", "./cache")
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/youtube/videos", handlers.FetchYouTubeVideos)
	app.Get("/post/:slug", handlers.GetSinglePost)
	app.Get("/posts", handlers.GetAllPosts)

}
