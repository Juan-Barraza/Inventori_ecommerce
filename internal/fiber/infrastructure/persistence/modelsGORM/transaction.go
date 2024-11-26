package modelsgorm

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Status   string
	Amount   float64
	Currency string
	Date     time.Time
	OrderId  uint
	Order    Order `gorm:"foreignKey:OrderId"`
}

type TransactionJson struct {
	ID       uint
	Status   string
	Amount   float64
	Currency string
	Date     time.Time
	OrderId  uint
}
