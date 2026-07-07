package reservations

import "parkora/internal/reservations/dto"

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
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

func (s *service) GetAllReservations() (*dto.GetReservationResponse, error) {
	reservations, err := s.repo.GetAllReservations()
	if err != nil {
		return nil, err
	}

	responses := make([]*dto.ReservationResponse, 0, len(reservations))

	for _, reservation := range reservations {
		responses = append(responses, reservation.ToResponse())
	}

	return &dto.GetReservationResponse{
		Success: true,
		Message: "Reservations retrieved successfully",
		Data:    responses,
	}, nil
}
