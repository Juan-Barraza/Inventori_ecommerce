package ports

import (
	"context"
	"time"
)

type IPaymentGateway interface {
	IPayment
	CreateOrder(ctx context.Context, amount float64, currency, returnURL, cancelURL string) (orderID, approvalURL string, err error)
	CaptureOrder(orderID string) (status string, capturedAt time.Time, err error)
	VerifyWebhook(body []byte, headers map[string]string) (bool, error)
}
