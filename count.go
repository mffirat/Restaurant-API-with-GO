package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CountHandler(c *fiber.Ctx) error {

	var floor1, floor2, floor3 int
	var total int

	for i := 1; i <= 3; i++ {
		key := "floor:" + strconv.Itoa(i)
		valStr, err := client.Get(ctx, key).Result()
		var say int
		if err != nil || valStr == "" {
			var dbCount int64
			DB.Model(&Customer{}).Where("floor = ? AND exited_at IS NULL", i).Count(&dbCount)
			say = int(dbCount)
		} else {
			say, err = strconv.Atoi(valStr)
			if err != nil {
				return c.JSON(fiber.Map{"error": "Invalid say"})
			}
		}

		if i == 1 {
			floor1 = say
		} else if i == 2 {
			floor2 = say
		} else {
			floor3 = say
		}
	}

	total = floor1 + floor2 + floor3

	response := FloorCount{
		Floor1: floor1,
		Floor2: floor2,
		Floor3: floor3,
		Total:  total,
	}

	return c.JSON(response)
}
