package ports

import "context"

type PaymentRequest struct {
	Amount    float64
	Currency  string
	Method    string
	ReturnURL *string
	CancelURL *string
}

type PaymentResponse struct {
	Status        string
	ReturnURL     *string
	CancelURL     *string
	TransactionID *string
}

type IPayment interface {
	Process(ctx context.Context, req PaymentRequest) (PaymentResponse, error)
}
