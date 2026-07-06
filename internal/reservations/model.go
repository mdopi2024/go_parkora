package reservations

import (
	"parkora/internal/parking"
	"parkora/internal/users"

	"gorm.io/gorm"
)

// ReservationStatus represents the status of a reservation.
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
	LicensePlate string            `gorm:"type:varchar(15);not null" json:"license_plate"`
	Status       ReservationStatus `gorm:"type:varchar(20);default:active;not null" json:"status"`

	User users.User          `gorm:"foreignKey:UserID" json:"user"`
	Zone parking.ParkingZone `gorm:"foreignKey:ZoneID" json:"zone"`
}
