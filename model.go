package main

import "time"

type Customer struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Gender    string     `json:"gender"`
	AgeGroup  string     `json:"age_group"`
	Floor     int        `json:"floor"`
	Payment   float64    `json:"payment"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	ExitedAt  *time.Time `json:"exited_at"`
}

type FloorCount struct {
	Floor1 int `json:"floor1"`
	Floor2 int `json:"floor2"`
	Floor3 int `json:"floor3"`
	Total  int `json:"total"`
}
