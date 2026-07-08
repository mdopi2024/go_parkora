package reservations

import (
	"errors"

	"parkora/internal/parking"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	CreateReservation(reservation *Reservation) error
	GetAllReservations() ([]Reservation, error)
	GetReservationByID(id uint) (*Reservation, error)
	GetReservationByIDAndUserID(reservationID uint, userID uint) (*Reservation, error)
	UpdateReservationFields(req *Reservation) error
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
	ErrParkingZoneNotFound              = errors.New("parking zone not found")
	ErrParkingZoneFull                  = errors.New("parking zone is full")
	ErrReservationNotFound              = errors.New("reservation not found")
	ErrForbiddenAccess                  = errors.New("you are not able to update this reservation")
	ErrFailedToActions                  = errors.New("Faile oparations")
	ErrCannotModifyCancelledReservation = errors.New("cancelled reservations cannot be modified")
	ErrCannotChangeCompletedReservation = errors.New("completed reservation status cannot be changed")
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

func (r *repository) GetAllReservations() ([]Reservation, error) {
	var reservations []Reservation

	if err := r.db.
		Preload("Zone").
		Find(&reservations).Error; err != nil {
		return nil, err
	}

	return reservations, nil
}

func (r *repository) GetReservationByID(id uint) (*Reservation, error) {
	var reservation Reservation

	err := r.db.
		Preload("Zone").
		First(&reservation, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrReservationNotFound
		}
		return nil, err
	}

	return &reservation, nil
}
func (r *repository) GetReservationByIDAndUserID(reservationID uint, userID uint) (*Reservation, error) {
	var reservation Reservation

	err := r.db.
		Where("id = ? AND user_id = ?", reservationID, userID).
		Preload("Zone").
		First(&reservation).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrReservationNotFound
		}
		return nil, err
	}

	return &reservation, nil
}

func (r *repository) UpdateReservationFields(reservation *Reservation) error {
	return r.db.Save(reservation).Error
}
