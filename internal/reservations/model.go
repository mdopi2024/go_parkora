package reservations

import (
	"parkora/internal/parking"
	"parkora/internal/reservations/dto"
	"parkora/internal/users"

	"gorm.io/gorm"
)

type ReservationStatus string

const (
	ReservationStatusActive    ReservationStatus = "active"
	ReservationStatusCompleted ReservationStatus = "completed"
	ReservationStatusCancelled ReservationStatus = "cancelled"
)

type Reservation struct {
	gorm.Model

	UserID       uint              `gorm:"not null;index" json:"user_id"`
	ZoneID       uint              `gorm:"not null;index" json:"zone_id"`
	LicensePlate string            `gorm:"size:15;not null" json:"license_plate"`
	Status       ReservationStatus `gorm:"size:20;default:'active';not null;check:status IN ('active','completed','cancelled')" json:"status"`

	User users.User          `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Zone parking.ParkingZone `gorm:"foreignKey:ZoneID" json:"zone,omitempty"`
}

func (r *Reservation) ToResponse() *dto.ReservationResponse {
	return &dto.ReservationResponse{
		ID:           r.ID,
		UserID:       r.UserID,
		ZoneID:       r.ZoneID,
		LicensePlate: r.LicensePlate,
		Status:       string(r.Status),
		CreatedAt:    r.CreatedAt,
		UpdatedAt:    r.UpdatedAt,
	}

}
