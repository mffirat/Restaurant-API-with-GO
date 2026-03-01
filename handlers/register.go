package handlers

import (
	"Go2/domain"

	"github.com/gofiber/fiber/v2"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// @Summary Register new user
// @Description Create a new user account
// @Tags auth
// @Accept json
// @Produce json
// @Param data body handlers.RegisterRequest true "Register payload"
// @Success 201 {object} map[string]string "User registered successfully"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 500 {object} map[string]string "User could not be created"
// @Router /register [post]
func RegisterHandler(c *fiber.Ctx, service *domain.DomainService) error {
	var rq RegisterRequest
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
	if err := service.RegisterUser(rq.Username, rq.Password); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "user could not be created",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered succesfully!",
	})
}
