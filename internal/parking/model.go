package parking

import (
	parkingdto "parkora/internal/parking/dto"

	"gorm.io/gorm"
)

type ParkingType string

const (
	ParkingTypeCar   ParkingType = "car"
	ParkingTypeBike  ParkingType = "bike"
	ParkingTypeTruck ParkingType = "truck"
	ParkingTypeBus   ParkingType = "bus"
)

type ParkingZone struct {
	gorm.Model

	Name          string      `gorm:"size:100;not null" json:"name"`
	Type          ParkingType `gorm:"size:20;not null;check:type IN ('car','bike','truck','bus')" json:"type"`
	TotalCapacity int         `gorm:"not null;check:total_capacity > 0" json:"total_capacity"`
	PricePerHour  float64     `gorm:"type:decimal(10,2);not null;check:price_per_hour > 0" json:"price_per_hour"`
}

func (p *ParkingZone) ToResponse() parkingdto.ParkingResponse {
	return parkingdto.ParkingResponse{
		ID:            p.ID,
		Name:          p.Name,
		Type:          string(p.Type),
		TotalCapacity: p.TotalCapacity,
		PricePerHour:  p.PricePerHour,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
	}
}
