package commands

import (
	"fmt"
	"inventory/internal/fiber/domain/ports"
	"inventory/internal/fiber/domain/repositories"
	"log"
)

type CaptureOrderPaypal struct {
	txRepo        repositories.ITransactionRepository
	gateWayPaypal ports.IPaymentGateway
}

func NewCaptureOrderPaypal(txRepo repositories.ITransactionRepository,
	gateWayPaypal ports.IPaymentGateway) *CaptureOrderPaypal {
	return &CaptureOrderPaypal{
		txRepo:        txRepo,
		gateWayPaypal: gateWayPaypal,
	}
}

func (c *CaptureOrderPaypal) CaptureOrderPaypal(orderId uint) (string, error) {
	tx, err := c.txRepo.FindByOrderID(orderId)
	if err != nil {
		return "", fmt.Errorf("error to get transaction")
	}

	if tx == nil {
		return "", fmt.Errorf("transaction not found")
	}

	status, _, err := c.gateWayPaypal.CaptureOrder(*tx.TransactionId)
	if err != nil {
		return "", fmt.Errorf("error to get order paypal %w", err)
	}

	log.Println("Status del capturar orderPaypal: %w", status)

	if status == "COMPLETED" && tx.Status == "PENDING" {
		tx.Status = status
		c.txRepo.Update(tx)
	}

	return status, nil
}
