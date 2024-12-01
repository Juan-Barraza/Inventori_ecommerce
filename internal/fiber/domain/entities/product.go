package domain

type Product struct {
	Name        string
	Description string
	Price       float64
	Images      []PicturesProduct
	Stock       int
	CategoryId  uint
	ProviderId  uint
}
