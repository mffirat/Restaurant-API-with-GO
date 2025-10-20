package main

import (
	"log"

	"Go2/model"
	"Go2/repository/postgresql"
	"Go2/repository/redis"

	"github.com/gofiber/fiber/v2"

	redisClient "github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=postgres user=postgres password=1234 dbname=mydb port=5432 sslmode=disable TimeZone=Asia/Istanbul"
	var err error
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	DB.AutoMigrate(&model.Customer{})

	client := redisClient.NewClient(&redisClient.Options{
		Addr: "redis:6379",
	})

	customerRepo := postgresql.NewCustomerRepo(DB)
	floorRepo := redis.NewFloorRepo(client)

	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {

		return UpdateHandler(c, customerRepo, floorRepo)
	})
	app.Get("/count", func(c *fiber.Ctx) error {
		return CountHandler(c, floorRepo)
	})
	app.Get("/total_customers", func(c *fiber.Ctx) error {
		return TotalCustomersHandler(c, customerRepo)
	})
	app.Get("/children", func(c *fiber.Ctx) error {
		return ChildrenHandler(c, customerRepo)
	})
	app.Get("/total_income", func(c *fiber.Ctx) error {
		return TotalIncomeHandler(c, customerRepo)
	})

	log.Fatal(app.Listen(":8000"))

}
