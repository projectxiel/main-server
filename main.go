package main

import (
	"log"
	"projectxiel/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	// app.Use(cache.New(cache.Config{Expiration: time.Hour * 24}))
	routes.Routes(app)
	log.Fatal(app.Listen(":6969"))
}
