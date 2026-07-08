package server

import (
	"net/http"
	"parkora/internal/config"
	"parkora/internal/parking"
	"parkora/internal/reservations"
	"parkora/internal/users"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally return the error to let each route control the status code.
		return echo.ErrBadRequest.Wrap(err)
	}
	return nil
}

func StartServer(db *gorm.DB) {
	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Parkora API!")
	})

	users.AuthRegister(e, db)
	parking.RegisterRoutes(e, db)
	reservations.ReservationRoutes(e, db)

	if err := e.Start(":" + config.LoadEnv().Port); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
