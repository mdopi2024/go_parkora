package parking

import "gorm.io/gorm"

type Repository interface {
	CreateParkingZone(zone *ParkingZone) error
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
