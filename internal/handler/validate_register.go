package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/ihsanpraditya/gin-clean-1/graph/model"
)

func ValidateRegisterInput(validate *validator.Validate, input model.RegisterInput) error {
	if err := validate.Struct(input); err != nil {
		return err
	}

	if input.Password != input.ConfirmPassword {
		return CustomFieldError{
			Field:   "ConfirmPassword",
			Message: "Confirm password must match the password",
		}
	}

	return nil
}