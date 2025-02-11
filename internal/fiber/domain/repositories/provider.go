package repositories

import (
	domain "inventory/internal/fiber/domain/entities"

	"gorm.io/gorm"
)

type IProviderRepository interface {
	CreateProvider(provider *domain.Provider) error
	GetAllProvider() (*gorm.DB, error)
	GetById(id uint) (*domain.Provider, error)
	Update(provider *domain.Provider) error
	Delete(provider *domain.Provider) error
	GetByPhoneNummber(number string) (*domain.Provider, error)
}
