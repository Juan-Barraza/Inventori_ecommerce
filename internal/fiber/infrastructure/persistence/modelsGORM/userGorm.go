package modelsgorm

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

type UserGormJson struct {
	ID       uint
	Email    string
	Password string
}
