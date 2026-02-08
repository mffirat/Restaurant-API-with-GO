package main

import (
	"Go2/domain"
	userModel "Go2/domain/user"
	"Go2/handlers"
	"Go2/model"

	"Go2/repository/postgresql"
	"Go2/repository/redis"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"

	redisClient "github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"Go2/middlewares"
	"github.com/joho/godotenv"
)


func main() {
if err := godotenv.Load(); err != nil {
	log.Println("No .env file found")
}




	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")

	
	if host == "" || user == "" || password == "" || dbname == "" || port == "" {
		log.Fatal("One or more required environment variables are not set")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Istanbul",
		host, user, password, dbname, port,
	)

	var err error
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	DB.AutoMigrate(&model.Customer{})
	DB.AutoMigrate(&userModel.User{})

	client := redisClient.NewClient(&redisClient.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})

	customerRepo := postgresql.NewCustomerRepo(DB)
	floorRepo := redis.NewFloorRepo(client)
	userRepo := postgresql.NewUserRepo(DB)

	service := domain.NewDomainService(customerRepo, floorRepo, userRepo)

	app := fiber.New()
	app.Use(middlewares.RequestBodyLog)

	app.Post("/", func(c *fiber.Ctx) error {

		return handlers.UpdateHandler(c, service)
	})
	app.Post("/register", func(c *fiber.Ctx) error {
		return handlers.RegisterHandler(c, service)
	})
	app.Post("/login", func(c *fiber.Ctx) error {
		return handlers.LoginHandler(c, service)
	})
	app.Get("/count", func(c *fiber.Ctx) error {
		return handlers.CountHandler(c, service)
	})
	app.Get("/total_customers",middlewares.JWTAuth(), func(c *fiber.Ctx) error {
		return handlers.TotalCustomersHandler(c, service)
	})
	app.Get("/children",middlewares.JWTAuth(), func(c *fiber.Ctx) error {
		return handlers.ChildrenHandler(c, service)
	})
	app.Get("/total_income",middlewares.OnlyAdmin(), func(c *fiber.Ctx) error {
		return handlers.TotalIncomeHandler(c, service)
	})

	log.Fatal(app.Listen(":8000"))

}
