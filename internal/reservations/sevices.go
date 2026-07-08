package reservations

import (
	"parkora/internal/middleware"
	"parkora/internal/reservations/dto"
)

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

func (s *service) GetReservationByID(id uint) (*dto.GetReservationByIDResponse, error) {
	reservation, err := s.repo.GetReservationByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.GetReservationByIDResponse{
		Success: true,
		Message: "Reservation retrieved successfully",
		Data:    reservation.ToResponse(),
	}, nil
}

func (s *service) UpdateReservationFields(
	id uint,
	userID uint,
	role string,
	req *dto.UpdateReservationRequest,
) (*dto.ReservationResponse, error) {

	var (
		reservation *Reservation
		err         error
	)

	// Admin can update any reservation
	// Driver can update only their own reservation
	if role == string(middleware.RoleAdmin) {
		reservation, err = s.repo.GetReservationByID(id)
	} else {
		reservation, err = s.repo.GetReservationByIDAndUserID(id, userID)
	}

	if err != nil {
		return nil, err
	}

	// Driver cannot modify cancelled reservation
	if role == string(middleware.RoleDriver) &&
		reservation.Status == ReservationStatusCancelled {

		return nil, ErrCannotModifyCancelledReservation
	}

	// Completed reservation status cannot be changed
	if reservation.Status == ReservationStatusCompleted &&
		req.Status != "" {

		return nil, ErrCannotChangeCompletedReservation
	}

	// Update License Plate
	if req.LicensePlate != "" {
		reservation.LicensePlate = req.LicensePlate
	}

	// Update Zone
	if req.ZoneID != nil {
		reservation.ZoneID = *req.ZoneID
	}

	// Update Status
	if req.Status != "" {
		reservation.Status = ReservationStatus(req.Status)
	}

	if err := s.repo.UpdateReservationFields(reservation); err != nil {
		return nil, err
	}

	return reservation.ToResponse(), nil
}
