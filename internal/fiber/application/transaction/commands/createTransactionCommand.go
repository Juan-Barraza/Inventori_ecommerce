package commands

import (
	"context"
	// "crypto/rand"
	"fmt"
	paymentfactory "inventory/internal/fiber/application/paymentFactory"
	domain "inventory/internal/fiber/domain/entities"
	"inventory/internal/fiber/domain/ports"
	"inventory/internal/fiber/domain/repositories"
	"inventory/internal/fiber/infrastructure/paidMethod/paypal"

	// "math/big"
	"time"
)

type CreateTransactionCommand struct {
	txRepo         repositories.ITransactionRepository
	factoryPayment *paymentfactory.PaymentFactory
}

func NewCreateTransactionCommand(
	txRepo repositories.ITransactionRepository,
	factoryPayment *paymentfactory.PaymentFactory,

) *CreateTransactionCommand {
	return &CreateTransactionCommand{txRepo: txRepo, factoryPayment: factoryPayment}
}

func (c *CreateTransactionCommand) ProcessPaymentTransaction(ctx context.Context, orderID uint, dto ports.PaymentRequest) (string, string, error) {
	extingTx, err := c.txRepo.FindByOrderID(orderID)
	if err != nil {
		return "", "", fmt.Errorf("error checking for existing transaction: %w", err)
	}

	if extingTx != nil {
		return extingTx.Status, "", nil
	}

	provider, err := c.factoryPayment.GetProviderPayment(dto.Method)
	if err != nil {
		return "", "", fmt.Errorf("error to get provider to payment:  %w", err)
	}
	// guardar en BD
	tx := &domain.Transaction{
		Status:   "CREATED",
		Amount:   dto.Amount,
		Currency: dto.Currency,
		Date:     time.Now(),
		OrderId:  orderID,
	}

	if err := c.txRepo.Create(tx); err != nil {
		return "", "", fmt.Errorf("failed to create transaction for OrderID %d: %w", orderID, err)
	}

	resp, err := provider.Process(ctx, dto)
	if err != nil {
		tx.Status = "FAILED"
		c.txRepo.Update(tx)
		return "", "", fmt.Errorf("error to process pay %w", err)
	}

	if _, ok := provider.(*paypal.PayPalAdapter); ok {
		if resp.TransactionID == nil {
			return "", "", fmt.Errorf("PayPal did not return a valid transaction ID")
		}
		tx.TransactionId = resp.TransactionID
		tx.Status = "PENDING"
		c.txRepo.Update(tx)
	}

	return tx.Status, *resp.ReturnURL, nil

}

// func generateUniqueTransactionID() string {
// 	bigInt, err := rand.Int(rand.Reader, big.NewInt(1000000000))
// 	if err != nil {
// 		return ""
// 	}
// 	return fmt.Sprintf("TX-%d-%d", time.Now().Unix(), bigInt)
// }
