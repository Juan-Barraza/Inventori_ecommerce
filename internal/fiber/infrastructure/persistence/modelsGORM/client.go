package modelsgorm

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Name           string `gorm:"not null"`
	LastName       string `gorm:"not null"`
	TypeDocument   string
	DocumentNumber string `gorm:"unique;not null"`
	PhoneNumber    string `gorm:"not null"`
	Address        string
	UserID         uint
	User           User `gorm:"foreignKey:UserID"`
}

type ClientJson struct {
	ID             uint
	Name           string
	LastName       string
	TypeDocument   string
	Email          string
	Password       string
	DocumentNumber string
	PhoneNumber    string
	Address        string
	UserId         uint
}
