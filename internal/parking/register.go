package parking

import (
	"parkora/internal/auth"
	"parkora/internal/config"
	"parkora/internal/middleware"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	env := config.LoadEnv()
	repository := NewRepository(db)
	jwtService := auth.NewJWTService(env.SecretKey)
	service := NewService(repository)
	handler := NewHandler(service)

	api := e.Group("/api/v1/zones")
	api.POST("", handler.Create, middleware.RoleBasedAuthMiddleware(jwtService, middleware.RoleAdmin))
	api.GET("", handler.GetAll, middleware.RoleBasedAuthMiddleware(jwtService, middleware.RoleAdmin, middleware.RoleDriver))
	api.GET("/:id", handler.GetByID, middleware.RoleBasedAuthMiddleware(jwtService, middleware.RoleAdmin, middleware.RoleDriver))
	api.PUT("/:id", handler.Update, middleware.RoleBasedAuthMiddleware(jwtService, middleware.RoleAdmin))

}
