package validators

import (
	"fmt"
	domain "inventory/internal/fiber/domain/entities"
)

func ValidateOrder(order *domain.OrderDTO) error {
	if order.Status == "" {
		return fmt.Errorf("status is required")
	}
	if order.Quantity <= 0 {
		return fmt.Errorf("quantity must be greater than 0")
	}
	if order.Date.IsZero() {
		return fmt.Errorf("date is required")
	}
	if order.ClientId == 0 {
		return fmt.Errorf("client ID is required")
	}
	if len(order.Products) == 0 {
		return fmt.Errorf("at least one product is required")
	}
	
	return nil
}