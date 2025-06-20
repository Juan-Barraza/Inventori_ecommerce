package repositories

import domain "inventory/internal/fiber/domain/entities"

type ITransactionRepository interface {
	Create(transaction *domain.Transaction) error
	GetById(id uint) (*domain.Transaction, error)
	FindByOrderID(orderID uint) (*domain.Transaction, error)
	Update(tx *domain.Transaction) error
	GetByTransactionID(transactionID string) (*domain.Transaction, error)
}
