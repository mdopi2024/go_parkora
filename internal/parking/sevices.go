package parking

import (
	"errors"

	httpresponse "parkora/internal/httpResponse"
	parkingdto "parkora/internal/parking/dto"

	"gorm.io/gorm"
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

func (s *service) GetAllParkingZones() (*parkingdto.GetAllParkingZonesResponse, *httpresponse.ErrorResponse) {
	zones, err := s.repo.GetAllParkingZones()
	if err != nil {
		return nil, &httpresponse.ErrorResponse{
			Success: false,
			Message: "failed to retrieve parking zones",
			Errors:  err.Error(),
		}
	}

	data := make([]parkingdto.ParkingResponse, 0, len(zones))
	for _, z := range zones {
		data = append(data, parkingdto.ParkingResponse{
			ID:             z.ID,
			Name:           z.Name,
			Type:           z.Type,
			TotalCapacity:  z.TotalCapacity,
			AvailableSpots: z.AvailableSpots,
			PricePerHour:   z.PricePerHour,
		})
	}

	return &parkingdto.GetAllParkingZonesResponse{
		Success: true,
		Message: "Parking zones retrieved successfully",
		Data:    data,
	}, nil
}

func (s *service) GetParkingZoneByID(id uint) (*parkingdto.GetParkingZoneResponse, *httpresponse.ErrorResponse) {
	zone, err := s.repo.GetParkingZoneByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &httpresponse.ErrorResponse{
				Success: false,
				Message: "parking zone not found",
			}
		}
		return nil, &httpresponse.ErrorResponse{
			Success: false,
			Message: "failed to retrieve parking zone",
			Errors:  err.Error(),
		}
	}

	return &parkingdto.GetParkingZoneResponse{
		Success: true,
		Message: "Parking zone retrieved successfully",
		Data: parkingdto.ParkingResponse{
			ID:             zone.ID,
			Name:           zone.Name,
			Type:           zone.Type,
			TotalCapacity:  zone.TotalCapacity,
			AvailableSpots: zone.AvailableSpots,
			PricePerHour:   zone.PricePerHour,
		},
	}, nil
}

func (s *service) Update(id uint, req *parkingdto.UpdateParkingRequest) (*parkingdto.UpdateParkingResponse, *httpresponse.ErrorResponse) {
	// Fetch the existing record from the database
	zone, err := s.repo.GetParkingZoneModel(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &httpresponse.ErrorResponse{
				Success: false,
				Message: "parking zone not found",
			}
		}
		return nil, &httpresponse.ErrorResponse{
			Success: false,
			Message: "failed to retrieve parking zone",
			Errors:  err.Error(),
		}
	}

	// Only overwrite fields that are non-empty / non-zero; keep existing values otherwise
	if req.Name != "" {
		zone.Name = req.Name
	}
	if req.Type != "" {
		zone.Type = ParkingType(req.Type)
	}
	if req.TotalCapacity > 0 {
		zone.TotalCapacity = req.TotalCapacity
	}
	if req.PricePerHour > 0 {
		zone.PricePerHour = req.PricePerHour
	}

	if err := s.repo.SaveParkingZone(zone); err != nil {
		return nil, &httpresponse.ErrorResponse{
			Success: false,
			Message: "failed to update parking zone",
			Errors:  err.Error(),
		}
	}

	return &parkingdto.UpdateParkingResponse{
		Success: true,
		Message: "Parking zone updated successfully",
		Data:    zone.ToResponse(),
	}, nil
}


