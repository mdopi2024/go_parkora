package middleware

import (
	"fmt"
	"net/http"
	"slices"
	"strings"

	"parkora/internal/auth"
	httpresponse "parkora/internal/httpResponse"

	"github.com/labstack/echo/v5"
)

const (
	AuthorizationHeader = "Authorization"
	BearerScheme        = "Bearer"
	ClaimsContextKey    = "claims"
)

type Role string

const (
	RoleAdmin  Role = "admin"
	RoleDriver Role = "driver"
)

// RoleBasedAuthMiddleware(jwtService, RoleAdmin, RoleDriver)
func RoleBasedAuthMiddleware(
	jwtService auth.JWTService,
	allowedRoles ...Role,
) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c *echo.Context) error {

			// Get Authorization header
			authHeader := c.Request().Header.Get(AuthorizationHeader)
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, httpresponse.ErrorResponse{
					Success: false,
					Message: "authorization header is required",
				})
			}

			// Validate Bearer token format
			parts := strings.Fields(authHeader)
			if len(parts) != 2 || !strings.EqualFold(parts[0], BearerScheme) {
				return c.JSON(http.StatusUnauthorized, httpresponse.ErrorResponse{
					Success: false,
					Message: "authorization header must be in the format: Bearer <token>",
				})
			}

			// Validate JWT
			claims, err := jwtService.ValidateToken(parts[1])
			if err != nil {
				// Log err internally if needed.
				fmt.Println("JWT Error:", err)
				return c.JSON(http.StatusUnauthorized, httpresponse.ErrorResponse{
					Success: false,
					Message: "invalid or expired token",
				})
			}

			// Store claims in context for handlers
			c.Set(ClaimsContextKey, claims)

			// If no roles are specified, any authenticated user is allowed.
			if len(allowedRoles) == 0 {
				return next(c)
			}

			// Check whether the user's role is allowed.
			userRole := Role(claims.Role)

			if !slices.Contains(allowedRoles, userRole) {
				return c.JSON(http.StatusForbidden, httpresponse.ErrorResponse{
					Success: false,
					Message: "you do not have permission to access this resource",
				})
			}

			return next(c)
		}
	}
}

// GetClaims returns the JWT claims stored in the context.
func GetClaims(c *echo.Context) *auth.Claims {
	claims, ok := c.Get(ClaimsContextKey).(*auth.Claims)
	if !ok {
		return nil
	}

	return claims
}
