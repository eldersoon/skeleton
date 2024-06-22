package controllers

import "github.com/gofiber/fiber/v2"

func HelloController(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"Message": "Hello, u!",
	})
}