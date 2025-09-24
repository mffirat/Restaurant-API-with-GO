package main

import (
	"github.com/gofiber/fiber/v2"
)

func TotalCustomersHandler(c *fiber.Ctx) error {
	startDate := c.Query("start")
	endDate := c.Query("end")

	var count int64
	if err := DB.Model(&Customer{}).
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Count(&count).Error; err != nil {
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"total_customers": count})
}

func ChildrenHandler(c *fiber.Ctx) error {
	startDate := c.Query("start")
	endDate := c.Query("end")

	var count int64
	if err := DB.Model(&Customer{}).
		Where("age_group = ? AND created_at BETWEEN ? AND ?", "child", startDate, endDate).
		Count(&count).Error; err != nil {
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"children_count": count})
}

func TotalIncomeHandler(c *fiber.Ctx) error {
	startDate := c.Query("start")
	endDate := c.Query("end")

	var totalIncome float64
	if err := DB.Model(&Customer{}).
		Select("COALESCE(SUM(payment), 0)").
		Where("exited_at BETWEEN ? AND ?", startDate, endDate).
		Scan(&totalIncome).Error; err != nil {
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"total_income": totalIncome})
}
