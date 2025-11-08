package handlers

import (
	"Go2/model"
	
	"Go2/interfaces"

	"github.com/gofiber/fiber/v2"
)

func CountHandler(c *fiber.Ctx, redisRepo interfaces.FloorRepoInterface) error {

	floor1, err := redisRepo.GetFloorCount(1)
	if err != nil {
		return c.JSON(fiber.Map{"error": "Invalid say"})
	}
	floor2, err := redisRepo.GetFloorCount(2)
	if err != nil {
		return c.JSON(fiber.Map{"error": "Invalid say"})
	}
	floor3, err := redisRepo.GetFloorCount(3)
	if err != nil {
		return c.JSON(fiber.Map{"error": "Invalid say"})
	}
	
	total := floor1 + floor2 + floor3

	response := model.FloorCount{
		Floor1: floor1,
		Floor2: floor2,
		Floor3: floor3,
		Total:  total,
	}

	return c.JSON(response)
}
