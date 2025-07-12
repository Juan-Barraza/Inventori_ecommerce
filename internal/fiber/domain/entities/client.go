package domain

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
	User           User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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

type ClientResponse struct {
	ID             uint
	Name           string
	LastName       string
	TypeDocument   string
	Email          string
	DocumentNumber string
	PhoneNumber    string
	Address        string
	UserId         uint
}

func ToClient(c *Client) *ClientResponse {
	return &ClientResponse{
		ID:             c.ID,
		Name:           c.Name,
		LastName:       c.LastName,
		TypeDocument:   c.TypeDocument,
		Email:          c.User.Email,
		DocumentNumber: c.DocumentNumber,
		PhoneNumber:    c.PhoneNumber,
		Address:        c.Address,
		UserId:         c.UserID,
	}
}
