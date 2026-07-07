package reservations

import (
	"errors"
	"net/http"

	httpresponse "parkora/internal/httpResponse"
	"parkora/internal/middleware"
	"parkora/internal/reservations/dto"

	"github.com/labstack/echo/v5"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateReservation(c *echo.Context) error {
	var req dto.CreateReservationRequest

	// Bind request body
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, &httpresponse.ErrorResponse{
			Success: false,
			Message: "invalid request body",
		})
	}

	// Validate request body
	if err := c.Validate(&req); err != nil {
		return httpresponse.ValidationError(c, err)
	}

	// Get authenticated user ID from JWT middleware
	userID := middleware.GetClaims(c).UserID
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, &httpresponse.ErrorResponse{
			Success: false,
			Message: "unauthorized",
		})
	}

	resp, err := h.service.CreateReservation(userID, req)
	if err != nil {
		switch {
		case errors.Is(err, ErrParkingZoneNotFound):
			return c.JSON(http.StatusNotFound, &httpresponse.ErrorResponse{
				Success: false,
				Message: err.Error(),
			})

		case errors.Is(err, ErrParkingZoneFull):
			return c.JSON(http.StatusConflict, &httpresponse.ErrorResponse{
				Success: false,
				Message: err.Error(),
			})

		default:
			return c.JSON(http.StatusInternalServerError, &httpresponse.ErrorResponse{
				Success: false,
				Message: "failed to create reservation",
			})
		}
	}

	return c.JSON(http.StatusCreated, resp)
}
