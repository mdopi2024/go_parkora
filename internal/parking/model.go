package parking

import (
	"time"

	"gorm.io/gorm"
)

type ParkingZone struct {
	gorm.Model

	Name          string    `gorm:"type:varchar(100);not null" json:"name"`
	Type          string    `gorm:"type:varchar(30);not null" json:"type"`
	TotalCapacity int       `gorm:"type:int;not null;check:total_capacity > 0" json:"total_capacity"`
	PricePerHour  float64   `gorm:"type:decimal(10,2);not null;check:price_per_hour > 0" json:"price_per_hour"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
