package mappers

import (
	domain "inventory/internal/fiber/domain/models"
	modelsgorm "inventory/internal/fiber/infrastructure/persistence/modelsGORM"
)

func ToOrderGorm(o *domain.Order) *modelsgorm.Order {
	return &modelsgorm.Order{
		Status:      o.Status,
		Quantity:    o.Quantity,
		Date:        o.Date,
		Description: o.Description,
		ClientId:    o.ClientId,
	}
}

func FromOrderGorm(o *modelsgorm.Order) *domain.Order {
	return &domain.Order{
		Status:      o.Status,
		Quantity:    o.Quantity,
		Date:        o.Date,
		Description: o.Description,
		ClientId:    o.ClientId,
	}
}
