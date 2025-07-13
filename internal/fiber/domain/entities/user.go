package domain

import "gorm.io/gorm"

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

type UserResponse struct {
	ID    uint
	Email string
}

func ToUser(u *User) *UserResponse {
	return &UserResponse{
		ID:    u.ID,
		Email: u.Email,
	}
}
