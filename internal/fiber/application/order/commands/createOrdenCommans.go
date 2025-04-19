package commands

import (
	"fmt"
	domain "inventory/internal/fiber/domain/entities"
	"inventory/internal/fiber/domain/repositories"
)

type CreateOrderCommand struct {
	orderRepo repositories.IOrderRepository
}

func NewCreateOrderCommand(orderRepo repositories.IOrderRepository) *CreateOrderCommand {
	return &CreateOrderCommand{orderRepo: orderRepo}
}

func (s *CreateOrderCommand) CreateOrder(orderData *domain.Order) error {
	if err := s.orderRepo.CreateOrder(orderData); err != nil {
		return fmt.Errorf("error creating order: %w", err)
	}
	return nil
}
