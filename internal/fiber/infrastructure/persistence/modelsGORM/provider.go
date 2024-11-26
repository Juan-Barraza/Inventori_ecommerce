package modelsgorm

import "gorm.io/gorm"

type Provider struct {
	gorm.Model
	Name          string
	Address       string
	PhoneNumber   string
	TypeOfProduct string
	UserID        uint
	User          User `gorm:"foreignKey:UserID"`
}

type ProviderJson struct {
	ID            uint
	Name          string
	Address       string
	PhoneNumber   string
	TypeOfProduct string
	UserID        uint
}
