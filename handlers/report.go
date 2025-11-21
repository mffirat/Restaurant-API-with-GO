package handlers

import (
	
	"Go2/interfaces"

	"github.com/gofiber/fiber/v2"
)

func TotalCustomersHandler(c *fiber.Ctx, repo interfaces.CustomerRepoInterface) error {
	startDate := c.Query("start")
	endDate := c.Query("end")

	count, err := repo.GetTotalCustomers(startDate, endDate)
	if err != nil {
		return c.JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.JSON(fiber.Map{"total_customers": count})
}

func ChildrenHandler(c *fiber.Ctx, repo interfaces.CustomerRepoInterface) error {
	startDate := c.Query("start")
	endDate := c.Query("end")

	count, err := repo.GetChildrenCount(startDate, endDate)
	if err != nil {
		return c.JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.JSON(fiber.Map{"children_count": count})
}

func TotalIncomeHandler(c *fiber.Ctx, repo interfaces.CustomerRepoInterface) error {
	startDate := c.Query("start")
	endDate := c.Query("end")

	total, err := repo.GetTotalIncome(startDate, endDate)
	if err != nil {
		return c.JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.JSON(fiber.Map{"total_income": total})
}
