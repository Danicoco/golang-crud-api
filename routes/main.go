package routes

import "github.com/gofiber/fiber/v2"

func Routes(app *fiber.App) {
	v1 := app.Group("/v1", func(c *fiber.Ctx) error {
		return c.Next()
	})

	BookRoutes(v1)

	v1.Get("/", func(c *fiber.Ctx) error {
		return c.Send([]byte("Welcome to Book API"))
	})
}
