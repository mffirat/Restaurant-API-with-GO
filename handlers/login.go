package handlers

import (
	"Go2/domain"

	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// @Summary Login
// @Description Login with username and password then returns JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param data body handlers.LoginRequest true "Login payload"
// @Success 200 {object} map[string]string "JWT token"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 401 {object} map[string]string "Invalid username or password"
// @Router /login [post]
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
	token, err := service.LoginUser(rq.Username, rq.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid username or password",
		})
	}
	return c.JSON(fiber.Map{
		"token": token,
	})
}
