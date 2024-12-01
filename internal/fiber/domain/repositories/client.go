package repositories

import domain "inventory/internal/fiber/domain/entities"

type IClientRepository interface {
	Create(client *domain.Client) error
	GetAll() ([]domain.Client, error)
	GetById(id uint) (*domain.Client, error)
	Update(client *domain.Client) error
	Delete(client *domain.Client) error
	FindByDocumentNumber(docNumber string) (*domain.Client, error)
}
