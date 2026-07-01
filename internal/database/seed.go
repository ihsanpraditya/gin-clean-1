package database

import (
	"log"
	"github.com/ihsanpraditya/gin-clean-1/internal/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	var dbRoles []model.Role
	var count int64

	log.Println("Checking for existing users in the database...")
	db.Model(&model.User{}).Count(&count)

	if (count > 0) {
		log.Println("Users already exist in the database. Skipping seeding.")
		return
	}

	log.Println("Running database seeders...")

	// 1. Ensure Roles exist (Fallback check)
	roles := []string{"super_admin", "admin", "user"}
	for _, roleName := range roles {
		var role model.Role
		err := db.FirstOrCreate(&role, model.Role{Name: roleName}).Error
		if err != nil {
			log.Printf("Failed to seed role %s: %v", roleName, err)
		}
		dbRoles = append(dbRoles, role)
	}

	// Seed a Default Super Admin if none exists
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)

	// Create the super_admin user
	superAdmin := model.User{
		Name:     "Super Admin",
		Email:    "superadmin@example.com",
		Password: string(hashedPassword),
	}

	if err := db.Create(&superAdmin).Error; err == nil {
		// Find the super_admin role object to associate
		var adminRole model.Role
		db.Where("name = ?", "super_admin").First(&adminRole)

		// GORM Many-to-Many association insert
		db.Model(&superAdmin).Association("Roles").Append(&adminRole)
		log.Println("Default Super Admin successfully seeded (superadmin@example.com / password123)")
	} else {
		log.Printf("Failed to seed Super Admin: %v", err)
	}
}
