package graph

import (
	"github.com/go-playground/validator/v10"
	"github.com/ihsanpraditya/gin-clean-1/internal/service"
)

type Resolver struct {
	UserSvc *service.UserService
	Validator *validator.Validate
}
