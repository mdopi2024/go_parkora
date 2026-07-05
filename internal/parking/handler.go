package parking

import (
	"net/http"

	httpresponse "parkora/internal/httpResponse"
	parkingdto "parkora/internal/parking/dto"

	"github.com/labstack/echo/v5"
)

type Handler struct {
	service *service
}

func NewHandler(service *service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(c *echo.Context) error {
	var req parkingdto.CreateParkingRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, &httpresponse.ErrorResponse{
			Success: false,
			Message: "invalid request payload",
			Errors:  err,
		})
	}

	if err := c.Validate(req); err != nil {
		return httpresponse.ValidationError(c, err)
	}

	resp, errResp := h.service.Create(&req)
	if errResp != nil {
		return c.JSON(http.StatusBadRequest, &httpresponse.ErrorResponse{
			Success: false,
			Message: errResp.Message,
			Errors:  errResp.Errors,
		})
	}

	return c.JSON(http.StatusCreated, resp)
}
