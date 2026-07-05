package users

import (
	"fmt"
	"net/http"
	"strings"

	httpresponse "parkora/internal/httpResponse"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
)

func ValidationError(c *echo.Context, err error) error {
	fieldErrors := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range validationErrors {
			fieldErrors[strings.ToLower(fe.Field())] = validationMessage(fe)
		}
	} else {
		fieldErrors["error"] = err.Error()
	}

	return c.JSON(http.StatusBadRequest, httpresponse.ErrorResponse{
		Success: false,
		Message: "Validation failed",
		Errors:  fieldErrors,
	})
}

func validationMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fe.Field())

	case "email":
		return "Email must be a valid email address"

	case "min":
		return fmt.Sprintf("%s must be at least %s characters", fe.Field(), fe.Param())

	case "max":
		return fmt.Sprintf("%s must not exceed %s characters", fe.Field(), fe.Param())

	default:
		return fe.Error()
	}
}
