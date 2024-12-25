package domain

type Provider struct {
	ID            uint
	Name          string
	Address       string
	PhoneNumber   string
	TypeOfProduct string
	UserID        uint
	User
}
