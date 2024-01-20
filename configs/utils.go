package configs

import "github.com/gofiber/fiber/v2"

func Success(c *fiber.Ctx, message string, payload interface{}) error {
	return c.JSON(fiber.Map{
		"data":    payload,
		"message": message,
	})
}

func Error(c *fiber.Ctx, message string, payload error) error {
	return c.Status(200).JSON(fiber.Map{
		"data":    payload,
		"message": message,
	})
}
