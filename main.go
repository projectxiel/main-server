package main

import (
	"log"
	"os"
	_ "projectxiel/docs"
	"projectxiel/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":6969"
	} else {
		port = ":" + port
	}

	return port
}
func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	// app.Use(cache.New(cache.Config{Expiration: time.Hour * 24}))
	routes.Routes(app)
	log.Fatal(app.Listen(getPort()))
}
