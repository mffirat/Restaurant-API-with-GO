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
	app.Post("/", UpdateHandler)
	app.Get("/count", CountHandler)
	app.Listen(":8080")

}
