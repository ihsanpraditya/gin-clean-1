package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/ihsanpraditya/gin-clean-1/internal/config"
)

var DB *gorm.DB

func ConnectDatabase(cfg config.DatabaseConfig) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Host,
        cfg.User,
        cfg.Password,
        cfg.Name,
        cfg.Port,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	log.Println("Connected to database.")
	DB = database

	Seed(DB)
}
