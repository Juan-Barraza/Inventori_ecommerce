package modelsgorm

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Status      string
	Quantity    int
	Date        time.Time
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
}
