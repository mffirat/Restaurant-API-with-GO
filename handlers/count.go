package handlers

import (
	"Go2/domain"

	"github.com/gofiber/fiber/v2"
)

// @Summary Get customer count per floor and total customers
// @Description Returns how many customers are on each floor and the total customer count
// @Tags Statistics
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /count [get]
func CountHandler(c *fiber.Ctx, service *domain.DomainService) error {

	counts, err := service.GetCounts()
	if err != nil {
		return c.JSON(fiber.Map{"error": "Floor counts couldn't get"})
	}
	return c.JSON(counts)
}
