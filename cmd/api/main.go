package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ihsanpraditya/gin-clean-1/internal/config"
	"github.com/ihsanpraditya/gin-clean-1/internal/database"
	"github.com/ihsanpraditya/gin-clean-1/internal/router"
)

func main() {
	cfg := config.LoadConfig()

	database.ConnectDatabase(cfg.DB)

	r := gin.Default()

	router.SetupRouter(r)

	r.Run(":8080")
}
