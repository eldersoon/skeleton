package routes

import (
	"github.com/eldersoon/go-csv-store/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// Hello Route!!!!!!!!!
	api := app.Group("/api")

	api.Get("/hello", controllers.HelloController)
}