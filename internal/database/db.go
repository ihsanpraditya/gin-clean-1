package database

import (
	"fmt"
	"log"
	"time"

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

	sqlDB, err := DB.DB()
	if err != nil {
		log.Printf("Gagal mengambil instance sql.DB: %v", err)
	} else {
		// Batasi maksimal proses Postgres yang boleh dibuat oleh aplikasi Gin Anda
		// Untuk 100 user bersamaan, angka 20-25 sudah sangat aman dan efisien
		sqlDB.SetMaxOpenConns(25)

		// Jumlah koneksi standby/idle yang tetap dibiarkan hidup di pool
		sqlDB.SetMaxIdleConns(10)

		// Durasi maksimal koneksi boleh digunakan sebelum otomatis dibuat baru oleh Go
		// Berguna untuk menghindari kebocoran memori pada koneksi yang terlalu lama
		sqlDB.SetConnMaxLifetime(1 * time.Hour)
		
		// Durasi maksimal koneksi idle boleh bertahan di pool sebelum ditutup otomatis
		sqlDB.SetConnMaxIdleTime(15 * time.Minute)

		log.Println("Database connection pool configured successfully.")
	}

	Seed(DB)
}
