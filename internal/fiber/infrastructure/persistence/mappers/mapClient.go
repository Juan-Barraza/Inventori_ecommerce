package mappers

import (
	domain "inventory/internal/fiber/domain/models"
	modelsgorm "inventory/internal/fiber/infrastructure/persistence/modelsGORM"
)

func ToClientGorm(c *domain.Client) *modelsgorm.Client {
	return &modelsgorm.Client{
		Name:           c.Name,
		LastName:       c.LastName,
		TypeDocument:   c.TypeDocument,
		DocumentNumber: c.DocumentNumber,
		PhoneNumber:    c.PhoneNumber,
		Address:        c.Address,
		UserID:         c.UserID,
	}
}

func FromlientGorm(c *modelsgorm.Client) *domain.Client {
	return &domain.Client{
		Name:           c.Name,
		LastName:       c.LastName,
		TypeDocument:   c.TypeDocument,
		DocumentNumber: c.DocumentNumber,
		PhoneNumber:    c.PhoneNumber,
		Address:        c.Address,
		UserID:         c.UserID,
	}
}
