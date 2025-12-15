package middlewares

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func RequestBodyLog(c *fiber.Ctx) error {
	body := c.Body()
	if len(body) == 0 {
		query := c.OriginalURL()
		log.Printf("Request Body: %s\n", query)
	} else {
		log.Println("Request Body: %s\n", string(body))
	}
	return c.Next()
}
