package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func UpdateHandler(c *fiber.Ctx) error {

	EntranceStr := c.Query("Entrance", "0")
	ExitStr := c.Query("Exit", "0")
	FloorStr := c.Query("Floor", "1")

	Entrance, err := strconv.Atoi(EntranceStr)
	if err != nil {
		return c.JSON(fiber.Map{"error": "Invalid value for Entrance: " + err.Error()})
	}
	Exit, err := strconv.Atoi(ExitStr)
	if err != nil {
		return c.JSON(fiber.Map{"error": "Invalid value for Exit: " + err.Error()})
	}
	Floor, err := strconv.Atoi(FloorStr)
	if err != nil || Floor < 1 || 3 < Floor {
		return c.JSON(fiber.Map{"error": "Invalid value for Floor: " + err.Error()})

	}
	key := "floor:" + strconv.Itoa(Floor)

	var change = Entrance - Exit

	if change > 0 {
		_, err = client.IncrBy(ctx, key, int64(change)).Result()
		if err != nil {
			return c.JSON(fiber.Map{"error": "Redis INCRBY error: " + err.Error()})
		}
	} else if change < 0 {
		_, err = client.DecrBy(ctx, key, int64(-change)).Result()
		if err != nil {
			return c.JSON(fiber.Map{"error": "Redis DECRBY error: " + err.Error()})
		}
	}
	response := UpdateResponse{
		Message: "Customer nums updated",
		Floor:   Floor,
		Current: change,
	}

	return c.JSON(response)
}
