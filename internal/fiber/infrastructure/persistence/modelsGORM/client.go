package modelsgorm

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Name           string
	LastName       string
	TypeDocument   string
	DocumentNumber string
	PhoneNumber    string
	Address        string
	UserID         uint
	User           User `gorm:"foreignKey:UserID"`
}

type ClientJson struct {
	ID             uint
	Name           string
	LastName       string
	TypeDocument   string
	DocumentNumber string
	PhoneNumber    string
	Address        string
	UserId         uint
}
