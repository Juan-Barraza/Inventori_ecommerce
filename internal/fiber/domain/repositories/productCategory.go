package repositories

import (
	domain "inventory/internal/fiber/domain/entities"

	"gorm.io/gorm"
)

type ICategory interface {
	AddCategory(catgeory *domain.Category) error
	GetAllCategories() (*gorm.DB, error)
	GetById(id uint) (*domain.Category, error)
	DeleteCategory(category *domain.Category) error
	ValidateCategory(name string) error
}
