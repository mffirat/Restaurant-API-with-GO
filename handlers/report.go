package handlers

import (
	"Go2/domain"

	"github.com/gofiber/fiber/v2"
)

// @Summary Get total customers by date range
// @Description Returns total customers count in given date range
// @Tags stats
// @Accept json
// @Produce json
// @Param start query string true "Start date (YYYY-MM-DD)"
// @Param end query string true "End date (YYYY-MM-DD)"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]string
// @Router /total_customers [get]
func TotalCustomersHandler(c *fiber.Ctx, service *domain.DomainService) error {
	startDate := c.Query("start")
	endDate := c.Query("end")

	count, err := service.GetTotalCustomers(startDate, endDate)
	if err != nil {
		return c.JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.JSON(fiber.Map{"total_customers": count})
}

// @Security BearerAuth
// @Summary Get total income by date range (Admin only)
// @Description Returns total income between given dates (requires admin JWT)
// @Tags admin
// @Accept json
// @Produce json
// @Param start query string true "Start date (YYYY-MM-DD)"
// @Param end query string true "End date (YYYY-MM-DD)"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Router /total_income [get]
func ChildrenHandler(c *fiber.Ctx, service *domain.DomainService) error {
	startDate := c.Query("start")
	endDate := c.Query("end")

	count, err := service.GetChildrenCount(startDate, endDate)
	if err != nil {
		return c.JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.JSON(fiber.Map{"children_count": count})
}

// @Summary Get children count by date range
// @Description Returns total children count between dates
// @Tags stats
// @Accept json
// @Produce json
// @Param start query string true "Start date (YYYY-MM-DD)"
// @Param end query string true "End date (YYYY-MM-DD)"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]string
// @Router /children [get]
func TotalIncomeHandler(c *fiber.Ctx, service *domain.DomainService) error {
	startDate := c.Query("start")
	endDate := c.Query("end")

	total, err := service.GetTotalIncome(startDate, endDate)
	if err != nil {
		return c.JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.JSON(fiber.Map{"total_income": total})
}
