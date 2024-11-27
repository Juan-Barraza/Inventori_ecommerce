package repositories

import domain "inventory/internal/fiber/domain/models"

type IClientRepository interface {
	Create(client *domain.Client) error
	GetAll() ([]domain.Client, error)
	GetById(id int) (*domain.Client, error)
	FindByDocumentNumber(docNumber string) (*domain.Client, error)
}
