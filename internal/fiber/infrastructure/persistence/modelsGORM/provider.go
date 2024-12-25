package modelsgorm

import "gorm.io/gorm"

type Provider struct {
	gorm.Model
	Name          string `gorm:"not null"`
	Address       string `gorm:"not null"`
	PhoneNumber   string `gorm:"unique;not null"`
	TypeOfProduct string `gorm:"not null"`
	UserID        uint
	User          User `gorm:"foreignKey:UserID"`
}

type ProviderJson struct {
	ID            uint
	Name          string
	Address       string
	Email         string
	Password      string
	PhoneNumber   string
	TypeOfProduct string
	UserID        uint
}
