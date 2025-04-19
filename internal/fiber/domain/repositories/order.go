package repositories

import (
	domain "inventory/internal/fiber/domain/entities"

	"gorm.io/gorm"
)

type IOrderRepository interface {
	CreateOrder(order *domain.Order) error
	GetOrders(clientId, productId uint) (*gorm.DB, []domain.Order, error)
	GetByID(id uint) (*domain.Order, error)
	UpdateOrder(order *domain.Order) error
	DeleteOrder(orderId uint) error
}
