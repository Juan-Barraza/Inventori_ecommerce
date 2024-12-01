package domain

type Client struct {
	ID             uint
	Name           string
	LastName       string
	TypeDocument   string
	DocumentNumber string
	PhoneNumber    string
	Address        string
	UserID         uint
	User
}
