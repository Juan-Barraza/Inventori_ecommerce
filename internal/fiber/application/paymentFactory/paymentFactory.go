package paymentfactory

import (
	"fmt"
	"inventory/internal/fiber/domain/ports"
	"inventory/internal/fiber/infrastructure/paidMethod/paypal"
)

type PaymentFactory struct {
	providers map[string]ports.IPayment
}

func NewPaymentFactory() *PaymentFactory {
	payPalAdapter, err := paypal.NewPayPalAdapter()
	if err != nil {
		fmt.Println("Error to create paypal adapter")
	}

	return &PaymentFactory{
		providers: map[string]ports.IPayment{
			"paypal": payPalAdapter,
		},
	}
}

func (p *PaymentFactory) GetProviderPayment(method string) (ports.IPayment, error) {
	if provider, ok := p.providers[method]; ok {
		return provider, nil
	}

	return nil, fmt.Errorf("Unknown payment: %s", method)
}
