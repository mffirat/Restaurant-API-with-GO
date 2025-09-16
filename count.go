package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CountHandler(c *fiber.Ctx) error {

	countStr1, err := client.Get(ctx, "floor:1").Result()
	if err != nil {
		return c.JSON(fiber.Map{"error": "Read error for Redis: " + err.Error()})
	}
	count1, err := strconv.Atoi(countStr1)
	if err != nil {
		return c.JSON(fiber.Map{"error": "Value conversation error for Redis: " + err.Error()})
	}
	countStr2, err := client.Get(ctx, "floor:2").Result()
	if err != nil {
		return c.JSON(fiber.Map{"error": "Read error for Redis: " + err.Error()})
	}
	count2, err := strconv.Atoi(countStr2)
	if err != nil {
		return c.JSON(fiber.Map{"error": "Value conversation error for Redis: " + err.Error()})
	}
	countStr3, err := client.Get(ctx, "floor:3").Result()
	if err != nil {
		return c.JSON(fiber.Map{"error": "Read error for Redis: " + err.Error()})
	}
	count3, err := strconv.Atoi(countStr3)
	if err != nil {
		return c.JSON(fiber.Map{"error": "Value conversation error for Redis: " + err.Error()})
	}

	response := FloorCount{
		Floor1: count1,
		Floor2: count2,
		Floor3: count3,
		Total:  count1 + count2 + count3,
	}

	return c.JSON(response)
}
