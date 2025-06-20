package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"inventory/internal/fiber/domain/ports"
	"inventory/internal/fiber/domain/repositories"
)

type WebHookService struct {
	txRepo        repositories.ITransactionRepository
	gatewayPaypal ports.IPaymentGateway
}

type PaypalWebHook struct {
	Status        string
	TransactionId string
}

func NewWebHookService(txRepo repositories.ITransactionRepository,
	gatewayPaypal ports.IPaymentGateway) *WebHookService {
	return &WebHookService{
		txRepo:        txRepo,
		gatewayPaypal: gatewayPaypal,
	}
}

func (w *WebHookService) HandleWebHook(ctx context.Context, body []byte, headers map[string]string) (string, error) {
	var webHookData PaypalWebHook
	if err := json.Unmarshal(body, &webHookData); err != nil {
		return "", fmt.Errorf("error unmarshaling webhook body: %w", err)
	}
	isValid, err := w.gatewayPaypal.VerifyWebhook(body, headers)
	if err != nil {
		return "", fmt.Errorf("error verifying webhook %w", err)
	}
	if !isValid {
		return "", fmt.Errorf("invalid webhook data")
	}
	tx, err := w.txRepo.GetByTransactionID(string(webHookData.TransactionId))
	if err != nil {
		return "", fmt.Errorf("error getting transaction by TransactionID: %w", err)
	}

	tx.Status = "COMPLETED"
	if err := w.txRepo.Update(tx); err != nil {
		return "", fmt.Errorf("error updating transaction status: %w", err)
	}

	return "Webhook processed successfully", nil
}
