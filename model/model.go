package model

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Gender    string     `json:"gender"`
	AgeGroup  string     `json:"age_group"`
	Floor     int        `json:"floor"`
	Payment   float64    `json:"payment"`
	EnteredAt time.Time  `json:"entered_at"`
	ExitedAt  *time.Time `json:"exited_at"`
}

type FloorCount struct {
	Floor1 int `json:"floor1"`
	Floor2 int `json:"floor2"`
	Floor3 int `json:"floor3"`
	Total  int `json:"total"`
}

