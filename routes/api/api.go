package api

import (
	"github.com/gofiber/fiber/v2"
	v1 "go_test/routes/api/v1"
)

func InitRoutes(app *fiber.App) {
	api := app.Group("/api", middleware)
	api.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"ping": "pong"})
	})
	v1.InitRoutes(api)
}

func middleware(c *fiber.Ctx) error {
	return c.Next()
}
