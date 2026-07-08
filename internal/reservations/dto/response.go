package dto

import (
	"time"
)

type ZoneResponse struct {
	ID            uint    `json:"id"`
	Name          string  `json:"name"`
	Type          string  `json:"type"`
	TotalCapacity int     `json:"total_capacity"`
	PricePerHour  float64 `json:"price_per_hour"`
}

type ReservationResponse struct {
	ID           uint         `json:"id"`
	UserID       uint         `json:"user_id"`
	ZoneID       uint         `json:"zone_id"`
	LicensePlate string       `json:"license_plate"`
	Status       string       `json:"status"`
	Zone         ZoneResponse `json:"zone"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}

type SuccessReservationResponse struct {
	Success bool                `json:"success"`
	Message string              `json:"message"`
	Data    ReservationResponse `json:"data"`
}
type GetReservationResponse struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Data    []*ReservationResponse `json:"data"`
}
type GetReservationByIDResponse struct {
	Success bool                 `json:"success"`
	Message string               `json:"message"`
	Data    *ReservationResponse `json:"data"`
}
