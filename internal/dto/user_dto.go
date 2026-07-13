package dto

type Role struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Roles	[]Role `json:"roles"`
	IsActive bool   `json:"is_active"`
}

type CreateUser struct {
	Name            string `json:"name" validate:"required,max=50"`
	Email           string `json:"email" validate:"required,max=100,email"`
	Password        string `json:"password" validate:"required,min=8,max=255,containsany=012345678"`
	ConfirmPassword string `json:"confirm_password"`
}

type UpdateUser struct {
	Name  			string `json:"name" validate:"required,max=50"`
	Email 			string `json:"email" validate:"required,max=100,email"`
	Password        string `json:"password" validate:"min=8,max=255,containsany=012345678"`
	ConfirmPassword string `json:"confirm_password" validate:"eqfield=Password"`
	Roles 			*[]uint `json:"roles" validate:"required"`
	IsActive 		bool `json:"is_active" validate:"required"`
}