package reservations

import (
	"parkora/internal/auth"
	"parkora/internal/config"
	"parkora/internal/middleware"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func ReservationRoutes(e *echo.Echo, db *gorm.DB) {
	env := config.LoadEnv()
	// Dependency Injection
	repo := NewRepository(db)
	jwtService := auth.NewJWTService(env.SecretKey)
	service := NewService(repo)
	handler := NewHandler(service)

	// Register routes
	group := e.Group("/api/v1/reservations")

	group.POST("", handler.CreateReservation, middleware.RoleBasedAuthMiddleware(jwtService))
	group.GET("", handler.GetAllReservations, middleware.RoleBasedAuthMiddleware(jwtService, middleware.RoleAdmin))
}
