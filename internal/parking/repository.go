package parking

import (
	"gorm.io/gorm"
)

type Repository interface {
	CreateParkingZone(zone *ParkingZone) error
	GetAllParkingZones() ([]ParkingZoneWithAvailable, error)
	GetParkingZoneByID(id uint) (*ParkingZoneWithAvailable, error)
	GetParkingZoneModel(id uint) (*ParkingZone, error)
	SaveParkingZone(zone *ParkingZone) error
}

type ParkingZoneWithAvailable struct {
	ID             uint
	Name           string
	Type           string
	TotalCapacity  int
	AvailableSpots int
	PricePerHour   float64
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateParkingZone(zone *ParkingZone) error {
	return r.db.Create(zone).Error
}

// GetAllParkingZones retrieves all parking zones with dynamically calculated available spots
func (r *repository) GetAllParkingZones() ([]ParkingZoneWithAvailable, error) {
	var zones []ParkingZoneWithAvailable

	// Subquery to count active reservations for each parking zone
	activeReservationsSubquery := r.db.
		Table("reservations").
		Select("COUNT(*)").
		Where("reservations.zone_id = parking_zones.id").
		Where("reservations.status = ?", "active")

	err := r.db.
		Table("parking_zones").
		Select(`
			parking_zones.id,
			parking_zones.name,
			parking_zones.type,
			parking_zones.total_capacity,
			parking_zones.price_per_hour,
			(parking_zones.total_capacity - COALESCE((?), 0)) AS available_spots
		`, activeReservationsSubquery).
		Scan(&zones).Error

	return zones, err
}

// GetParkingZoneByID retrieves a single parking zone with its dynamically calculated available spots
func (r *repository) GetParkingZoneByID(id uint) (*ParkingZoneWithAvailable, error) {
	var zone ParkingZoneWithAvailable

	activeReservationsSubquery := r.db.
		Table("reservations").
		Select("COUNT(*)").
		Where("reservations.zone_id = parking_zones.id").
		Where("reservations.status = ?", "active")

	err := r.db.
		Table("parking_zones").
		Select(`
			parking_zones.id,
			parking_zones.name,
			parking_zones.type,
			parking_zones.total_capacity,
			parking_zones.price_per_hour,
			(parking_zones.total_capacity - COALESCE((?), 0)) AS available_spots
		`, activeReservationsSubquery).
		Where("parking_zones.id = ?", id).
		Where("parking_zones.deleted_at IS NULL").
		First(&zone).Error

	if err != nil {
		return nil, err
	}

	return &zone, nil
}

// GetParkingZoneModel fetches the raw ParkingZone GORM model by ID.
func (r *repository) GetParkingZoneModel(id uint) (*ParkingZone, error) {
	var zone ParkingZone
	if err := r.db.First(&zone, id).Error; err != nil {
		return nil, err
	}
	return &zone, nil
}

// SaveParkingZone persists an existing ParkingZone record.
func (r *repository) SaveParkingZone(zone *ParkingZone) error {
	return r.db.Save(zone).Error
}
