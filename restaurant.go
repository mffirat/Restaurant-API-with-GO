package main

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var client *redis.Client

type FloorCount struct {
	Floor1 int `json:"floor1"`
	Floor2 int `json:"floor2"`
	Floor3 int `json:"floor3"`
	Total  int `json:"total"`
}
type UpdateResponse struct {
	Message string `json:"message"`
	Floor   int    `json:"floor"`
	Current int    `json:"current"`
}

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
	app.Post("/", func(c *fiber.Ctx) error {

		EntranceStr := c.Query("Entrance", "0")
		ExitStr := c.Query("Exit", "0")
		FloorStr := c.Query("Floor", "0")

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
	})

	app.Get("/count", func(c *fiber.Ctx) error {
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
	})
	app.Listen(":8080")

}
