package main

import (
	"github.com/eldersoon/go-csv-store/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	routes.Setup(app)

	app.Listen(":5000")
}