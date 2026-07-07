package reservations

import (
	"errors"

	"parkora/internal/parking"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	CreateReservation(reservation *Reservation) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

var (
	ErrParkingZoneNotFound = errors.New("parking zone not found")
	ErrParkingZoneFull     = errors.New("parking zone is full")
)

func (r *repository) CreateReservation(reservation *Reservation) error {
	return r.db.Transaction(func(tx *gorm.DB) error {

		// Lock the parking zone row
		var zone parking.ParkingZone

		if err := tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&zone, reservation.ZoneID).Error; err != nil {

			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrParkingZoneNotFound
			}

			return err
		}

		// Count active reservations
		var activeCount int64

		if err := tx.
			Model(&Reservation{}).
			Where("zone_id = ?", reservation.ZoneID).
			Where("status = ?", ReservationStatusActive).
			Count(&activeCount).Error; err != nil {
			return err
		}

		// Check parking capacity
		if activeCount >= int64(zone.TotalCapacity) {
			return ErrParkingZoneFull
		}

		// Create reservation
		if err := tx.Create(reservation).Error; err != nil {
			return err
		}

		return nil
	})
}
