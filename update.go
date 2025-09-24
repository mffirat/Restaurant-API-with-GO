package main

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func UpdateHandler(c *fiber.Ctx) error {

	action := c.Query("action", "enter")
	FloorStr := c.Query("Floor", "1")
	Gender := c.Query("Gender", "unknown")
	AgeGroup := c.Query("AgeGroup", "adult")
	idStr := c.Query("id", "0")
	paymentStr := c.Query("Payment", "0")
	Floor, err := strconv.Atoi(FloorStr)
	if err != nil || Floor < 1 || 3 < Floor {
		return c.JSON(fiber.Map{"error": "Invalid value for Floor: " + err.Error()})

	}
	key := "floor:" + strconv.Itoa(Floor)

	if action == "enter" {

		customer := Customer{
			Gender:    Gender,
			AgeGroup:  AgeGroup,
			Floor:     Floor,
			Payment:   0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			ExitedAt:  nil,
		}
		DB.Create(&customer)

		_, err := client.IncrBy(ctx, key, 1).Result()
		if err != nil {
			return c.JSON(fiber.Map{"error": "Redis INCRBY error: " + err.Error()})
		}

		return c.JSON(fiber.Map{
			"message": "Customer entered",
			"id":      customer.ID,
			"floor":   Floor,
		})
	} else if action == "exit" {

		id, err := strconv.Atoi(idStr)
		if err != nil || id <= 0 {
			return c.JSON(fiber.Map{"error": "Invalid ID"})
		}
		payment, err := strconv.ParseFloat(paymentStr, 64)
		if err != nil || id <= 0 {
			return c.JSON(fiber.Map{"error": "Invalid Payment"})
		}

		var customer Customer
		result := DB.First(&customer, id)
		if result.Error != nil {
			return c.JSON(fiber.Map{"error": "Customer not found"})
		}

		now := time.Now()
		customer.Payment = payment
		customer.ExitedAt = &now
		DB.Save(&customer)

		_, err = client.DecrBy(ctx, key, 1).Result()
		if err != nil {
			return c.JSON(fiber.Map{"error": "Redis DECRBY error: " + err.Error()})
		}

		return c.JSON(fiber.Map{
			"message":  "Customer exited",
			"id":       id,
			"payment":  payment,
			"exitedAt": now,
		})
	}

	return c.JSON(fiber.Map{"error": "Invalid action"})
}
