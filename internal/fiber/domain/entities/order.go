package domain

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Status      string    `gorm:"not null"`
	Quantity    int       `gorm:"not null"`
	Date        time.Time `gorm:"not null"`
	Description string
	ClientId    uint
	Client      Client    `gorm:"foreignKey:ClientId"`
	Products    []Product `gorm:"many2many:order_products"`
}

type OrderJson struct {
	ID          uint
	Status      string
	Quantity    int
	Date        time.Time
	Description string
	ClientId    uint
	Products    []ProductJson
}

type OrderDTO struct {
	ID          uint
	Status      string
	Quantity    int
	Date        time.Time
	Description string
	ClientId    uint
	Products    []uint
}

func ToOrder(p *Order) *OrderJson {
	prods := make([]ProductJson, len(p.Products))
	for i, prod := range p.Products {
		prods[i] = *ToProduct(&prod)
	}
	return &OrderJson{
		ID:          p.ID,
		Status:      p.Status,
		Quantity:    p.Quantity,
		Date:        p.Date,
		Description: p.Description,
		ClientId:    p.ClientId,
		Products:    prods,
	}
}
