package modelsgorm

import (
	"image"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
	Images      []PicturesProduct `gorm:"foreignKey:ProductId"`
	Stock       int
	CategoryId  uint
	Category    Category `gorm:"foreignKey:CategoryId"`
	ProviderId  uint
	Provider    Provider `gorm:"foreignKey:ProviderId"`
	Orders      []Order  `gorm:"many2many:order_products"`
}

type ProductJson struct {
	ID          uint
	Name        string
	Description string
	Price       float64
	Image       image.Image
	Stock       int
	CategoryId  uint
	ProviderId  uint
}
