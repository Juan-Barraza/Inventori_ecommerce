package repository

import (
	domain "inventory/internal/fiber/domain/entities"
	"inventory/pkg"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *pkg.Database
}

func NewOrderRepository(db *pkg.Database) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateOrder(order *domain.Order) error {
	return r.db.DB.Create(order).Error
}

func (r *OrderRepository) GetByID(id uint) (*domain.Order, error) {
	var order *domain.Order
	if err := r.db.DB.Preload("Products").First(&order, id).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (r *OrderRepository) GetOrders(clientId, productId uint) (*gorm.DB, []domain.Order, error) {
	var orders []domain.Order
	query := r.db.DB.Model(&domain.Order{}).
		Preload("Products.Images").
		Preload("Products").
		Preload("Client").
		Order("orders.id ASC")
	if clientId != 0 {
		query = query.Where("orders.client_id = ?", clientId)
	}

	if productId != 0 {
		query = query.Joins("JOIN order_products ON order_products.order_id = orders.id").
			Where("order_products.product_id = ?", productId)
	}

	result := query.Find(&orders)
	if result.Error != nil {
		return nil, nil, result.Error
	}

	return query, orders, nil
}

func (r *OrderRepository) UpdateOrder(order *domain.Order) error {
	return r.db.DB.Model(&order).Where("id = ?", order.ID).Updates(order).Error
}

func (r *OrderRepository) DeleteOrder(orderId uint) error {
	tx := r.db.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Unscoped().Where("order_id = ?", orderId).Delete(&domain.OrderProduct{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Delete(&domain.Order{}, orderId).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
