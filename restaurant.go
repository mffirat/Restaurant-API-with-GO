package main

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var client *redis.Client

func main() {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	countStr, err := client.Get(ctx, "currentCount").Result()
	if err != nil || countStr == "" {
		client.Set(ctx, "currentCount", 0, 0)
	}

	app := fiber.New()
	app.Post("/", func(c *fiber.Ctx) error {

		EntranceStr := c.Query("Entrance", "0")
		ExitStr := c.Query("Exit", "0")

		Entrance, err := strconv.Atoi(EntranceStr)
		if err != nil {
			return c.JSON(fiber.Map{"error": "Invalid value for Entrance"})
		}
		Exit, err := strconv.Atoi(ExitStr)
		if err != nil {
			return c.JSON(fiber.Map{"error": "Invalid value for Exit"})
		}

		var change = Entrance - Exit

		currentStr, err := client.Get(ctx, "currentCount").Result()
		if err != nil {
			return c.JSON(fiber.Map{"error": "Read error for Redis"})
		}
		current, err := strconv.Atoi(currentStr)
		if err != nil {
			return c.JSON(fiber.Map{"error": "Value conversation error for Redis"})
		}
		current += change
		if current < 0 {
			current = 0
		}
		err = client.Set(ctx, "currentCount", current, 0).Err()
		if err != nil {
			return c.JSON(fiber.Map{"error": "Write  error for Redis"})
		}
		return c.JSON(fiber.Map{
			"message ": "Customer nums updated", "current ": change,
		})
	})

	app.Get("/count", func(c *fiber.Ctx) error {
		countStr, err := client.Get(ctx, "currentCount").Result()
		if err != nil {
			return c.JSON(fiber.Map{"error": "Read error for Redis"})
		}
		count, err := strconv.Atoi(countStr)
		if err != nil {
			return c.JSON(fiber.Map{"error": "Value conversation error for Redis"})
		}

		return c.JSON(fiber.Map{
			"Count": count,
		})

	})
	app.Listen(":8080")

}
