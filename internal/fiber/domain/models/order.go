package domain

import "time"

type Order struct {
	Status      string
	Quantity    int
	Date        time.Time
	Description string
	Client
}
