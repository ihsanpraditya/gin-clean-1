package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/ihsanpraditya/gin-clean-1/internal/dto"
)

type MultiValidationError struct {
	ValidationErrors validator.ValidationErrors
	CustomErrors     []CustomFieldError
}

func (m MultiValidationError) Error() string {
	return "multiple validation errors occurred"
}

func ValidateRegisterInput(validate *validator.Validate, input *dto.CreateUser) error {
	var multiErr MultiValidationError
	
	if err := validate.Struct(input); err != nil {
		if valErrs, ok := err.(validator.ValidationErrors); ok {
			multiErr.ValidationErrors = valErrs
		}
	}

	if input.Password != input.ConfirmPassword {
		multiErr.CustomErrors = append(multiErr.CustomErrors, CustomFieldError{
			Field:   "ConfirmPassword",
			Message: "Confirm password must match the password",
		})
	}

	if len(multiErr.ValidationErrors) > 0 || len(multiErr.CustomErrors) > 0 {
		return multiErr
	}

	return nil
}