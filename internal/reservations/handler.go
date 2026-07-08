package reservations

import (
	"errors"
	"net/http"
	"strconv"

	httpresponse "parkora/internal/httpResponse"
	"parkora/internal/middleware"
	"parkora/internal/reservations/dto"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service *service
}

func NewHandler(service *service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) CreateReservation(c *echo.Context) error {
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

func (h *handler) GetAllReservations(c *echo.Context) error {
	resp, err := h.service.GetAllReservations()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &httpresponse.ErrorResponse{
			Success: false,
			Message: "failed to retrieve reservations",
		})
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *handler) GetReservationByID(c *echo.Context) error {
	// Get reservation ID from URL
	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil || id == 0 {
		return c.JSON(http.StatusBadRequest, &httpresponse.ErrorResponse{
			Success: false,
			Message: "invalid reservation id",
		})
	}

	// Call service
	resp, err := h.service.GetReservationByID(uint(id))
	if err != nil {

		switch {
		case errors.Is(err, ErrReservationNotFound):
			return c.JSON(http.StatusNotFound, &httpresponse.ErrorResponse{
				Success: false,
				Message: err.Error(),
			})

		default:
			return c.JSON(http.StatusInternalServerError, &httpresponse.ErrorResponse{
				Success: false,
				Message: "failed to retrieve reservation",
			})
		}
	}

	return c.JSON(http.StatusOK, resp)
}
