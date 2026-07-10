package handler

import (
	"errors"
	"fmt"
	"strings"
	"github.com/go-playground/validator/v10"
)

// ValidationError represents a single field error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// CustomFieldError allows us to pass manual field errors that mimic tag errors
type CustomFieldError struct {
	Field   string
	Message string
}

func (e CustomFieldError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

func FormatValidationErrors(err error) []ValidationError {
	var errorsList []ValidationError

	// 1. Check if it's our custom cross-field/manual validation error
	var customErr CustomFieldError
	if errors.As(err, &customErr) {
		return []ValidationError{{
			Field:   toSnakeCase(customErr.Field),
			Message: customErr.Message,
		}}
	}

	// 2. Check if it's a standard go-playground validation error slice
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return []ValidationError{{Field: "unknown", Message: err.Error()}}
	}

	for _, e := range validationErrors {
		errorsList = append(errorsList, ValidationError{
			Field:   toSnakeCase(e.Field()),
			Message: getErrorMessage(e),
		})
	}

	return errorsList
}

// getErrorMessage returns a human-readable message for each validation tag
func getErrorMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return fmt.Sprintf("Must be at least %s characters", e.Param())
	case "max":
		return fmt.Sprintf("Must be at most %s characters", e.Param())
	case "gte":
		return fmt.Sprintf("Must be greater than or equal to %s", e.Param())
	case "lte":
		return fmt.Sprintf("Must be less than or equal to %s", e.Param())
	case "oneof":
		return fmt.Sprintf("Must be one of: %s", e.Param())
	case "eqfield":
		return fmt.Sprintf("Must match %s", e.Param())
	case "phone":
		return "Invalid phone number format (use E.164: +1234567890)"
	case "slug":
		return "Must be a valid URL slug (lowercase letters, numbers, hyphens)"
	default:
		return fmt.Sprintf("Failed validation: %s", e.Tag())
	}
}

// toSnakeCase converts PascalCase to snake_case for JSON field names
func toSnakeCase(s string) string {
	var result []rune
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, r)
	}
	return strings.ToLower(string(result))
}
