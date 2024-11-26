package domain

import "time"

type Transaction struct {
	Status   string
	Amount   float64
	Currency string
	Date     time.Time
	Order
}
