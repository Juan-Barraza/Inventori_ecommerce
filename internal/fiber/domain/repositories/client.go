package repositories

import (
	domain "inventory/internal/fiber/domain/entities"

	"gorm.io/gorm"
)

type IClientRepository interface {
	Create(client *domain.Client) error
	GetAll() (*gorm.DB, error)
	GetById(id uint) (*domain.Client, error)
	Update(client *domain.Client) error
	Delete(client *domain.Client) error
	FindByDocumentNumber(docNumber string) (*domain.Client, error)
}
