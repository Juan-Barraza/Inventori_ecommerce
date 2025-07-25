package domain

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Status        string  `gorm:"not null"`
	Amount        float64 `gorm:"not null"`
	Currency      string  `gorm:"not null"`
	Date          time.Time
	OrderId       uint
	Order         Order `gorm:"foreignKey:OrderId"`
	TransactionId *string
}

type TransactionJson struct {
	ID            uint
	Status        string
	Amount        float64
	Currency      string
	Date          time.Time
	OrderId       uint
	TransactionID *string
}

func ToTransaction(t *Transaction) *TransactionJson {
	return &TransactionJson{
		ID:       t.ID,
		Status:   t.Status,
		Amount:   t.Amount,
		Currency: t.Currency,
		Date:     t.Date,
		OrderId:  t.OrderId,
	}
}
