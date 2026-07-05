package users

import (
	"net/http"

	httpresponse "parkora/internal/httpResponse"
	userdto "parkora/internal/users/dto"

	"github.com/labstack/echo/v5"
)

type Handler struct {
	service *UserService
}

func NewHandler(service *UserService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Register(c *echo.Context) error {
	var req userdto.RegisterUserRequest
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

	resp, errResp := h.service.Register(req)
	if errResp != nil {
		return c.JSON(http.StatusBadRequest, &httpresponse.ErrorResponse{
			Success: false,
			Message: errResp.Message,
			Errors:  errResp.Errors,
		})
	}

	return c.JSON(http.StatusCreated, resp)
}

func (h *Handler) Login(c *echo.Context) error {
	var req userdto.LoginUserRequest
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

	resp, errResp := h.service.Login(&req)
	if errResp != nil {
		return c.JSON(http.StatusUnauthorized, &httpresponse.ErrorResponse{
			Success: false,
			Message: errResp.Message,
			Errors:  errResp.Errors,
		})
	}

	return c.JSON(http.StatusOK, resp)
}
