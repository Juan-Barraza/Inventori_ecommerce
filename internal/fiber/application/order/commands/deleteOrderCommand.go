package commands

import (
	"fmt"
	"inventory/internal/fiber/domain/repositories"
)

type DeleteOrderCommand struct {
	orderRepo repositories.IOrderRepository
}

func NewDeleteOrderCommand(orderRepo repositories.IOrderRepository) *DeleteOrderCommand {
	return &DeleteOrderCommand{orderRepo: orderRepo}
}

func (d *DeleteOrderCommand) DeleteOrder(orderId uint) error {
	orderExists, err := d.orderRepo.GetByID(orderId)
	if err != nil {
		return fmt.Errorf("error checking order existence")
	}

	err = d.orderRepo.DeleteOrder(orderExists.ID)
	if err != nil {
		return fmt.Errorf("error deleting order")
	}

	return nil
}
