package handlers

import (
	
	
	"Go2/domain"

	"github.com/gofiber/fiber/v2"
)

func CountHandler(c *fiber.Ctx, service *domain.DomainService) error {

		counts, err := service.GetCounts()
	if err != nil {
		return c.JSON(fiber.Map{"error": "Floor counts couldn't get"})
	}
	return c.JSON(counts)
}