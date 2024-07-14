package routes

import "github.com/gofiber/fiber/v2"

func AllRoutes(app *fiber.App) {
	app.Get("/api/tes", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"Hello": "world"})
	})
}
