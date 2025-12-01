package handlers

import (
	
	"Go2/domain"

	"github.com/gofiber/fiber/v2"
)

func TotalCustomersHandler(c *fiber.Ctx, service *domain.DomainService) error {
	startDate := c.Query("start")
	endDate := c.Query("end")

	count, err := service.GetTotalCustomers(startDate, endDate)
	if err != nil {
		return c.JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.JSON(fiber.Map{"total_customers": count})
}

func ChildrenHandler(c *fiber.Ctx,  service *domain.DomainService) error {
	startDate := c.Query("start")
	endDate := c.Query("end")

	count, err := service.GetChildrenCount(startDate, endDate)
	if err != nil {
		return c.JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.JSON(fiber.Map{"children_count": count})
}

func TotalIncomeHandler(c *fiber.Ctx,  service *domain.DomainService) error {
	startDate := c.Query("start")
	endDate := c.Query("end")

	total, err := service.GetTotalIncome(startDate, endDate)
	if err != nil {
		return c.JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.JSON(fiber.Map{"total_income": total})
}
