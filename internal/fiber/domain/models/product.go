package domain

import "image"

type Product struct {
	Name        string
	Description string
	Price       float64
	Image       image.Image
	Stock       int
	Category
	Provider
}
