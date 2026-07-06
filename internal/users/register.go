package users

import (
	"parkora/internal/auth"
	"parkora/internal/config"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func AuthRegister(e *echo.Echo, db *gorm.DB) {
	env := config.LoadEnv()
	repository := NewRepository(db)
	jwtService := auth.NewJWTService(env.SecretKey)
	service := NewUserService(repository, jwtService)
	handler := NewHandler(service)
	api := e.Group("/api/v1/auth")
	api.POST("/register", handler.Register)
	api.POST("/login", handler.Login)
}
