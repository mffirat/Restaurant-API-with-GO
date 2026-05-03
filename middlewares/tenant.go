package middlewares
import (
    "github.com/gofiber/fiber/v2"
)

func TenantMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
        tenantID := c.Locals("tenant_id")
        if tenantID == nil {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": "tenant not found",
            })
        }
        return c.Next()
    }
}