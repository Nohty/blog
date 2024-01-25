package utils

import "github.com/gofiber/fiber/v2"

func Response(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"data": data,
	})
}
