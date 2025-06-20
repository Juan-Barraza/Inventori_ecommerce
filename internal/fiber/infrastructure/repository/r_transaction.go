package repository

import (
	domain "inventory/internal/fiber/domain/entities"
	"inventory/pkg"
)

type TransactionRepository struct {
	db *pkg.Database
}

func NewTransactionRepository(db *pkg.Database) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Create(transaction *domain.Transaction) error {
	return r.db.DB.Create(transaction).Error
}

func (r *TransactionRepository) GetById(id uint) (*domain.Transaction, error) {
	var transaction domain.Transaction
	if err := r.db.DB.First(&transaction, id).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *TransactionRepository) FindByOrderID(orderID uint) (*domain.Transaction, error) {
	var transaction *domain.Transaction
	if err := r.db.DB.Where(&domain.Transaction{OrderId: orderID}).First(&transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r *TransactionRepository) Update(tx *domain.Transaction) error {
	return r.db.DB.Model(&tx).Where("id = ?", tx.ID).Updates(tx).Error
}

func (r *TransactionRepository) GetByTransactionID(trID string) (*domain.Transaction, error) {
	var transaction domain.Transaction
	if err := r.db.DB.Model(&domain.Transaction{}).Where("transaction_id = ?", trID).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}
