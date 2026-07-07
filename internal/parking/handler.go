package parking

import (
	"net/http"
	"strconv"

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

func (h *Handler) GetAll(c *echo.Context) error {
	resp, errResp := h.service.GetAllParkingZones()
	if errResp != nil {
		return c.JSON(http.StatusInternalServerError, &httpresponse.ErrorResponse{
			Success: false,
			Message: errResp.Message,
			Errors:  errResp.Errors,
		})
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetByID(c *echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil || id == 0 {
		return c.JSON(http.StatusBadRequest, &httpresponse.ErrorResponse{
			Success: false,
			Message: "invalid zone id",
		})
	}

	resp, errResp := h.service.GetParkingZoneByID(uint(id))
	if errResp != nil {
		if errResp.Message == "parking zone not found" {
			return c.JSON(http.StatusNotFound, errResp)
		}
		return c.JSON(http.StatusInternalServerError, errResp)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) Update(c *echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil || id == 0 {
		return c.JSON(http.StatusBadRequest, &httpresponse.ErrorResponse{
			Success: false,
			Message: "invalid zone id",
		})
	}

	var req parkingdto.UpdateParkingRequest
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

	resp, errResp := h.service.Update(uint(id), &req)
	if errResp != nil {
		if errResp.Message == "parking zone not found" {
			return c.JSON(http.StatusNotFound, errResp)
		}
		return c.JSON(http.StatusBadRequest, errResp)
	}

	return c.JSON(http.StatusOK, resp)
}

