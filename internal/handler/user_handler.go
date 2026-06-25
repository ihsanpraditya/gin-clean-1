package handler

import (
	"net/http"
	"errors"
	"github.com/ihsanpraditya/docker-golang-postgres-api-boilerplate/internal/database"
	"github.com/ihsanpraditya/docker-golang-postgres-api-boilerplate/internal/model"
	"github.com/ihsanpraditya/docker-golang-postgres-api-boilerplate/internal/service"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

// Register handles user signups securely
func (h *UserHandler) Register(c *gin.Context) {
	var req RegisterRequest

	// 1. Validate payload structure and inputs
	if err := c.ShouldBindJSON(&req); err != nil {
		// You can customize this further via your internal/handler/validator.go
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Validation failed",
			"details": err.Error(),
		})
		return
	}

	// 2. Map DTO to business domain model safely
	userModel := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password, // Password hashing should happen downstream in service layer!
	}

	// 3. Execute business tier
	ctx := c.Request.Context()
	if err := h.svc.RegisterUser(ctx, userModel); err != nil {
		// 4. Inspect semantic domain errors gracefully
		if errors.Is(err, service.ErrEmailTaken) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "An unexpected error occurred"})
		return
	}

	// 5. Build clean, safe response payload (Zero exposure of Password hash)
	resp := UserResponse{
		ID:    userModel.ID, // Filled out automatically by GORM after successful creation
		Name:  userModel.Name,
		Email: userModel.Email,
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"data":    resp,
	})
}

// GET /users
func GetUsers(c *gin.Context) {
	var users []model.User
	database.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

// GET /users/:id
func GetUser(c *gin.Context) {
	var user model.User
	if err := database.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// PUT /users/:id
func UpdateUser(c *gin.Context) {
	var user model.User
	if err := database.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var input model.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&user).Updates(input)
	c.JSON(http.StatusOK, user)
}

// DELETE /users/:id
func DeleteUser(c *gin.Context) {
	var user model.User
	if err := database.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	database.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
