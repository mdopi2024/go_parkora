package parking

import (
	httpresponse "parkora/internal/httpResponse"
	parkingdto "parkora/internal/parking/dto"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) Create(req *parkingdto.CreateParkingRequest) (*parkingdto.CreateParkingResponse, *httpresponse.ErrorResponse) {
	zone := &ParkingZone{
		Name:          req.Name,
		Type:          ParkingType(req.Type),
		TotalCapacity: req.TotalCapacity,
		PricePerHour:  req.PricePerHour,
	}

	if err := s.repo.CreateParkingZone(zone); err != nil {
		return &parkingdto.CreateParkingResponse{}, &httpresponse.ErrorResponse{
			Success: false,
			Message: "failed to create parking zone",
			Errors:  err.Error(),
		}
	}

	return &parkingdto.CreateParkingResponse{
		Success: true,
		Message: "Parking zone created successfully",
		Data:    zone.ToResponse(),
	}, nil
}
