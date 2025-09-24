package main

import (
	"context"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ctx = context.Background()
var client *redis.Client
var DB *gorm.DB

func main() {
	dsn := "host=localhost user=postgres password=1234 dbname=mydb port=5432 sslmode=disable TimeZone=Asia/Istanbul"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	DB.AutoMigrate(&Customer{})

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
	app.Get("/total_customers", TotalCustomersHandler)
	app.Get("/children", ChildrenHandler)
	app.Get("/total_income", TotalIncomeHandler)
	app.Listen(":8080")

}
