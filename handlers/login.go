package handlers

import (
	"Go2/domain"

	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(c *fiber.Ctx, service *domain.DomainService) error {
	var rq LoginRequest
	if err := c.BodyParser(&rq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}
	if rq.Username == "" || rq.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "username or password Can not be empty",
		})
	}
	if err := service.LoginUser(rq.Username, rq.Password); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid username or password",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Login successful!",
	})
}
