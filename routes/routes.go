package routes

import (
	"projectxiel/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title ProjectXiel API
// @version 1.0
// @description Main API for the ProjectXiel website
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:6969
// @BasePath /
func Routes(app *fiber.App) {
	app.Get("/post/:slug", handlers.GetSinglePost)
	app.Get("/posts", handlers.GetAllPosts)
	app.Static("/", "./public")
	app.Get("/swagger/*", swagger.HandlerDefault)
}
