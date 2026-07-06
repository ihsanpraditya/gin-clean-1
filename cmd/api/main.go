package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/ihsanpraditya/gin-clean-1/internal/config"
	"github.com/ihsanpraditya/gin-clean-1/internal/database"
	"github.com/ihsanpraditya/gin-clean-1/internal/router"
)

func main() {
	cfg := config.LoadConfig()

	database.ConnectDatabase(cfg.DB)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:5173"}, // Add your frontend URLs here
			AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"}, // "Authorization" is required for your JWT middleware
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))

	router.SetupRouter(r)

	r.Run(":8080")
}
