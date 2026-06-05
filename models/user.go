package models

type User struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"title" binding:"required"`
	Email string `json:"author" binding:"required"`
	Password string `json:"password" binding:"required"`
}
