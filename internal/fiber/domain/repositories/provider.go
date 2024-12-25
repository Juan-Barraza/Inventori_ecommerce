package repositories

import domain "inventory/internal/fiber/domain/entities"

type IProviderRepository interface {
	CreateProvider(provider *domain.Provider) error
	GetAllProvider() ([]domain.Provider, error)
	GetById(id uint) (*domain.Provider, error)
	Update(provider *domain.Provider) error
	Delete(provider *domain.Provider) error
	GetByPhoneNummber(number string) (*domain.Provider, error)
}
