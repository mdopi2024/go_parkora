package reservations

import "parkora/internal/reservations/dto"

type Service interface {
	CreateReservation(userID uint, req dto.CreateReservationRequest) (*dto.SuccessReservationResponse, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateReservation(userID uint, req dto.CreateReservationRequest) (*dto.SuccessReservationResponse, error) {
	reservation := &Reservation{
		UserID:       userID,
		ZoneID:       req.ZoneID,
		LicensePlate: req.LicensePlate,
		Status:       ReservationStatusActive,
	}

	if err := s.repo.CreateReservation(reservation); err != nil {
		return nil, err
	}

	return &dto.SuccessReservationResponse{
		Success: true,
		Message: "Reservation confirmed successfully",
		Data:    *reservation.ToResponse(),
	}, nil
}
