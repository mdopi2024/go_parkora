package database

import (
	"parkora/internal/config"
	"parkora/internal/parking"
	"parkora/internal/reservations"
	"parkora/internal/users"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Pdb(env *config.Config) *gorm.DB {
	dsn := env.PG_DSN
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	migrateErr := Migrate(db)
	if migrateErr != nil {
		panic("failed to migrate database")
	}
	return db
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&users.User{},
		&parking.ParkingZone{},
		&reservations.Reservation{},
	)
}
