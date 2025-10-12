package main

import (
	"log"
	"strconv"
	"time"

	"Go2/model"
	"Go2/repository/postgresql"
	"Go2/repository/redis"

	"github.com/gofiber/fiber/v2"
)


func UpdateHandler(c *fiber.Ctx, repo postgresql.CustomerRepoInterface, redisRepo redis.FloorRepoInterface) error {

	action := c.Query("action", "enter")
	FloorStr := c.Query("Floor", "1")
	Gender := c.Query("Gender", "unknown")
	AgeGroup := c.Query("AgeGroup", "adult")
	idStr := c.Query("id", "0")
	paymentStr := c.Query("Payment", "0")
	Floor, err := strconv.Atoi(FloorStr)
	if err != nil || Floor < 1 || Floor > 3 {
		return c.JSON(fiber.Map{"error": "Invalid value for Floor: "})

	}
	

	if action == "enter" {

		customer := &model.Customer{
			Gender:    Gender,
			AgeGroup:  AgeGroup,
			Floor:     Floor,
			Payment:   0,
			EnteredAt: time.Now(),
			ExitedAt:  nil,
		}
		if err := repo.CreateCustomer(customer); err != nil {
			log.Println("Failed to create customer:", err)
			return c.JSON(fiber.Map{"error": "Internal Server Error"})
		}

		if err := redisRepo.IncreaseFloorCount(Floor); err != nil {
			return c.JSON(fiber.Map{"error": "Redis INCRBY error: "})
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
		if err != nil || payment <= 0 {
			return c.JSON(fiber.Map{"error": "Invalid Payment"})
		}

		customer, err := repo.GetCustomerByID(uint(id))
		if err != nil {
			return c.JSON(fiber.Map{"error": "Customer not found"})
		}

		now := time.Now()
		customer.Payment = payment
		customer.ExitedAt = &now
		if err := repo.UpdateCustomer(customer); err != nil {
			return c.JSON(fiber.Map{"error": "Internal Server Error"})
		}

		if err := redisRepo.DecreaseFloorCount(customer.Floor); err != nil {
			return c.JSON(fiber.Map{"error": "Redis DECRBY error: "})
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
