package mappers

import (
	domain "inventory/internal/fiber/domain/entities"
	modelsgorm "inventory/internal/fiber/infrastructure/persistence/modelsGORM"
)

func ToTransactionGorm(t *domain.Transaction) *modelsgorm.Transaction {
	return &modelsgorm.Transaction{
		Status:   t.Status,
		Amount:   t.Amount,
		Currency: t.Currency,
		Date:     t.Date,
		OrderId:  t.OrderId,
	}
}

func FromTransactionGorm(t *modelsgorm.Transaction) *domain.Transaction {
	return &domain.Transaction{
		Status:   t.Status,
		Amount:   t.Amount,
		Currency: t.Currency,
		Date:     t.Date,
		OrderId:  t.OrderId,
	}
}
