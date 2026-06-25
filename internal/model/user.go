package model

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string `json:"title" gorm:"type:varchar(50);not null"`
	Email     string `json:"author" gorm:"type:varchar(100);not null"`
	Password  string `json:"password" gorm:"type:varchar(255);not null"`
}
