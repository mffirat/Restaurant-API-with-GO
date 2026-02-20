package middlewares

import "github.com/gofiber/fiber/v2"

func OnlyAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := c.Locals("role")

		if role != "admin" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "forbidden",
			})
		}
		return c.Next()
	}
}
