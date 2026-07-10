package dto

type UserRegisterInput struct {
	Name            string `json:"name" validate:"required,max=50"`
	Email           string `json:"email" validate:"required,max=100,email"`
	Password        string `json:"password" validate:"required,min=8,max=255,containsany=012345678"`
	ConfirmPassword string `json:"confirmPassword"`
}