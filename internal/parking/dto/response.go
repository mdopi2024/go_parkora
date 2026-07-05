package dto

import "time"

type ParkingResponse struct {
	ID            uint      `json:"id"`
	Name          string    `json:"name"`
	Type          string    `json:"type"`
	TotalCapacity int       `json:"total_capacity"`
	PricePerHour  float64   `json:"price_per_hour"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type CreateParkingResponse struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Data    ParkingResponse `json:"data"`
}
