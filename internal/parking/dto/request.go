package dto

type CreateParkingRequest struct {
	Name          string  `json:"name" validate:"required"`
	Type          string  `json:"type" validate:"required"`
	TotalCapacity int     `json:"total_capacity" validate:"required,min=1"`
	PricePerHour  float64 `json:"price_per_hour" validate:"required,gt=0"`
}

type UpdateParkingRequest struct {
	Name          string  `json:"name" validate:"omitempty"`
	Type          string  `json:"type" validate:"omitempty"`
	TotalCapacity int     `json:"total_capacity" validate:"omitempty,min=1"`
	PricePerHour  float64 `json:"price_per_hour" validate:"omitempty,gt=0"`
}
