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

	for i := 1; i <= 3; i++ {
		key := "floor:" + strconv.Itoa(i)
		_, err := client.Get(ctx, key).Result()
		if err == redis.Nil {
			client.Set(ctx, key, 0, 0)
		}
	}

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {

		EntranceStr := c.Query("Entrance", "0")
		ExitStr := c.Query("Exit", "0")
		FloorStr := c.Query("Floor", "0")

		Entrance, err := strconv.Atoi(EntranceStr)
		if err != nil {
			return c.JSON(fiber.Map{"error": "Invalid value for Entrance"})
		}
		Exit, err := strconv.Atoi(ExitStr)
		if err != nil {
			return c.JSON(fiber.Map{"error": "Invalid value for Exit"})
		}
		Floor, err := strconv.Atoi(FloorStr)
		if err != nil || Floor < 1 || 3 < Floor {
			return c.JSON(fiber.Map{"error": "Invalid value for Floor"})

		}
		key := "floor:" + strconv.Itoa(Floor)

		var change = Entrance - Exit

		currentStr, err := client.Get(ctx, key).Result()
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
		err = client.Set(ctx, key, current, 0).Err()
		if err != nil {
			return c.JSON(fiber.Map{"error": "Write  error for Redis"})
		}
		return c.JSON(fiber.Map{
			"message ": "Customer nums updated",
			"floor":    Floor,
			"current ": change,
		})
	})

	app.Get("/count", func(c *fiber.Ctx) error {
		countStr1, err := client.Get(ctx, "floor:1").Result()
		if err != nil {
			return c.JSON(fiber.Map{"error": "Read error for Redis"})
		}
		count1, err := strconv.Atoi(countStr1)
		if err != nil {
			return c.JSON(fiber.Map{"error": "Value conversation error for Redis"})
		}
		countStr2, err := client.Get(ctx, "floor:2").Result()
		if err != nil {
			return c.JSON(fiber.Map{"error": "Read error for Redis"})
		}
		count2, err := strconv.Atoi(countStr2)
		if err != nil {
			return c.JSON(fiber.Map{"error": "Value conversation error for Redis"})
		}
		countStr3, err := client.Get(ctx, "floor:3").Result()
		if err != nil {
			return c.JSON(fiber.Map{"error": "Read error for Redis"})
		}
		count3, err := strconv.Atoi(countStr3)
		if err != nil {
			return c.JSON(fiber.Map{"error": "Value conversation error for Redis"})
		}
		return c.JSON(fiber.Map{
			"floor1": count1,
			"floor2": count2,
			"floor3": count3,
			"total":  count1 + count2 + count3,
		})
	})
	app.Listen(":8080")

}
