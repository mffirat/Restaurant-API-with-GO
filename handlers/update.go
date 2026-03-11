package handlers

import (
	"strconv"

	"Go2/domain"

	"github.com/gofiber/fiber/v2"
)

func UpdateHandler(c *fiber.Ctx, service *domain.DomainService) error {

	action := c.Query("action", "enter")
	FloorStr, err := strconv.Atoi(c.Query("Floor", "1"))
	if err != nil {
		return c.JSON(fiber.Map{"error": "Invalid floor number"})
	}
	Gender := c.Query("Gender", "unknown")
	AgeGroup := c.Query("AgeGroup", "adult")

	if action == "enter" {
		customer, _ := service.EnterCustomer(Gender, AgeGroup, FloorStr)

		return c.JSON(fiber.Map{
			"message": "Customer entered",
			"id":      customer.ID,
		})
	}

	if action == "exit" {
		id, _ := strconv.Atoi(c.Query("id", "0"))
		payment, _ := strconv.ParseFloat(c.Query("Payment", "0"), 64)
		if err := service.ExitCustomer(uint(id), payment); err != nil {
			return c.JSON(fiber.Map{"error": "could not exit"})
		}
		return c.JSON(fiber.Map{
			"message": "Customer exited",
			"id":      id,
			"payment": payment,
		})
	}

	return c.JSON(fiber.Map{"error": "Invalid action"})
}
