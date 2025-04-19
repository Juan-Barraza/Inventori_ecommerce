package domain

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `gorm:"index,not null"`
	Description string
	Price       float64           `gorm:"not null"`
	Images      []PicturesProduct `gorm:"foreignKey:ProductId;constraint:OnDelete:CASCADE;"`
	Stock       int               `gorm:"not null"`
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
	Image       []PicturesProduct
	Stock       int
	CategoryId  uint
	ProviderId  uint
}

func ToProduct(p *Product) *ProductJson {
	return &ProductJson{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Image:       p.Images,
		Stock:       p.Stock,
		CategoryId:  p.CategoryId,
		ProviderId:  p.ProviderId,
	}
}

func ToProductJson(p *ProductJson) *Product {
	return &Product{
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Images:      p.Image,
		Stock:       p.Stock,
		CategoryId:  p.CategoryId,
		ProviderId:  p.ProviderId,
	}
}
