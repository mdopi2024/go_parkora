package dto

type CreateReservationRequest struct {
	ZoneID       uint   `json:"zone_id" validate:"required"`
	LicensePlate string `json:"license_plate" validate:"required,max=15"`
}

type UpdateReservationRequest struct {
	ZoneID       *uint  `json:"zone_id" validate:"omitempty,gt=0"`
	LicensePlate string `json:"license_plate" validate:"omitempty,max=15"`
	Status       string `json:"status" validate:"omitempty,oneof=active cancelled completed"`
}
