package graph

import (
	"github.com/go-playground/validator/v10"
	"github.com/ihsanpraditya/gin-clean-1/internal/service"
)

type Resolver struct {
	RoleSvc *service.RoleService
	UserSvc *service.UserService
	AuthSvc *service.AuthService
	Validator *validator.Validate
}
