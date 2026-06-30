// Tambahkan struct ini di bagian atas bersama request DTO lainnya
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

// Tambahkan method ini di dalam user_handler.go
func (h *UserHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}

	ctx := c.Request.Context()
	token, err := h.svc.Login(ctx, req.Email, req.Password)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An unexpected error occurred"})
		return
	}

	// Ambil data user untuk disertakan di response (opsional namun umum)
	user, _ := h.svc.GetUserByEmail(ctx, req.Email)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"data": LoginResponse{
			Token: token,
			User: UserResponse{
				ID:    user.ID,
				Name:  user.Name,
				Email: user.Email,
			},
		},
	})
}
