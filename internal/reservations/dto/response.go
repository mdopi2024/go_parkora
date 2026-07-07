package dto

import (
	"time"
)

type ReservationResponse struct {
	ID           uint      `json:"id"`
	UserID       uint      `json:"user_id"`
	ZoneID       uint      `json:"zone_id"`
	LicensePlate string    `json:"license_plate"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
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
