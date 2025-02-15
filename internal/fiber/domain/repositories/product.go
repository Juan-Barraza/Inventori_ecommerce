package repositories

import (
	domain "inventory/internal/fiber/domain/entities"

	"gorm.io/gorm"
)

type IProduct interface {
	AddProduct(product *domain.Product) error
	GetAllProducts(category string, providerId uint) (*gorm.DB, []domain.Product, error)
	GetById(id uint) (*domain.Product, error)
	ValidateUniqueProduct(prod *domain.Product) error
	DeleteProduct(product *domain.Product) error
	UpdateProduct(product *domain.Product) error
}
