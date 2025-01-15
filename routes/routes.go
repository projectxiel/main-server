package routes

import (
	"projectxiel/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/projectxiel/fiberswagger"
)

func Routes(app *fiber.App) {
	app.Get("/swagger/*", fiberswagger.HandlerDefault)
	app.Get("/youtube/videos", handlers.FetchYouTubeVideos)
	app.Get("/post/:slug", handlers.GetSinglePost)
	app.Get("/posts", cache.New(), handlers.GetAllPosts)
	app.Get("/posts/search", handlers.SearchPosts)
	app.Get("/current-projects", cache.New(), handlers.GetCurrentProjects)
	app.Post("/update-cache", handlers.UpdateCache)
	//app.Get("/patreon/patrons", handlers.GetPatrons)
	app.Get("/:blobname", handlers.GetBlobData)
}
