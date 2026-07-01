package model

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"type:varchar(50);not null"`
	Email     string    `json:"email" gorm:"type:varchar(100);not null"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null"`
	Roles	  []Role    `json:"roles" gorm:"many2many:user_roles;"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
