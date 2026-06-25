package router

import (
	"github.com/ihsanpraditya/docker-golang-postgres-api-boilerplate/internal/database"
	"github.com/ihsanpraditya/docker-golang-postgres-api-boilerplate/internal/handler"
	"github.com/ihsanpraditya/docker-golang-postgres-api-boilerplate/internal/repository"
	"github.com/ihsanpraditya/docker-golang-postgres-api-boilerplate/internal/service"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	userRepo := repository.NewUserRepository(database.DB)
	userSvc := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userSvc)
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("", userHandler.Register)
		userRoutes.GET("", userHandler.GetUsers)
		userRoutes.GET("/:id", userHandler.GetUser)
		userRoutes.PUT("/:id", userHandler.UpdateUser)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}
}
