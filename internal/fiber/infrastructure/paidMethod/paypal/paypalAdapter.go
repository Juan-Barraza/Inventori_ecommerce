package paypal

import (
	"bytes"
	"context"
	"fmt"
	"inventory/internal/config"
	"inventory/internal/fiber/domain/ports"
	"io"
	"net/http"
	"time"

	"github.com/plutov/paypal/v4"
)

type PayPalAdapter struct {
	client *paypal.Client
}

func NewPayPalAdapter() (ports.IPaymentGateway, error) {
	c, err := paypal.NewClient(
		config.PayPalClientID,
		config.PayPalSecret,
		config.PayPalAPIBase,
	)
	if err != nil {
		return nil, err
	}
	if _, err := c.GetAccessToken(context.Background()); err != nil {
		return nil, fmt.Errorf("failed OAuth: %w", err)
	}
	return &PayPalAdapter{client: c}, nil
}

func (a *PayPalAdapter) Process(ctx context.Context, req ports.PaymentRequest) (ports.PaymentResponse, error) {
	OrderId, approveURL, err := a.CreateOrder(context.Background(), req.Amount, req.Currency, *req.ReturnURL, *req.CancelURL)
	if err != nil {
		return ports.PaymentResponse{}, nil
	}

	return ports.PaymentResponse{
		Status:        "CREATED",
		ReturnURL:     &approveURL,
		CancelURL:     req.CancelURL,
		TransactionID: &OrderId,
	}, nil
}

func (a *PayPalAdapter) CreateOrder(ctx context.Context, amount float64, currency, returnURL, cancelURL string) (string, string, error) {
	order, err := a.client.CreateOrder(
		ctx,
		"CAPTURE",
		[]paypal.PurchaseUnitRequest{{
			Amount: &paypal.PurchaseUnitAmount{
				Currency: currency,
				Value:    fmt.Sprintf("%.2f", amount),
			},
		}},
		nil, // no estás pasando información del payer
		&paypal.ApplicationContext{
			ReturnURL: returnURL,
			CancelURL: cancelURL,
		},
	)
	if err != nil {
		return "", "", err
	}

	// busca el link “approve”
	var approveURL string
	for _, l := range order.Links {
		if l.Rel == "approve" {
			approveURL = l.Href
			break
		}
	}
	if approveURL == "" {
		return order.ID, "", fmt.Errorf("no approval link found")
	}
	return order.ID, approveURL, nil
}

func (a *PayPalAdapter) CaptureOrder(orderID string) (string, time.Time, error) {
	resp, err := a.client.CaptureOrder(context.Background(), orderID, paypal.CaptureOrderRequest{})
	if err != nil {
		return "", time.Time{}, err
	}
	return string(resp.Status), time.Now(), nil
}

func (a *PayPalAdapter) VerifyWebhook(body []byte, headers map[string]string) (bool, error) {
	// Construimos un *http.Request para pasárselo al SDK
	req := &http.Request{
		Method: "POST",
		Header: http.Header{
			"Paypal-Transmission-Id":   []string{headers["paypal-transmission-id"]},
			"Paypal-Transmission-Time": []string{headers["paypal-transmission-time"]},
			"Paypal-Transmission-Sig":  []string{headers["paypal-transmission-sig"]},
			"Paypal-Cert-Url":          []string{headers["paypal-cert-url"]},
			"Paypal-Auth-Algo":         []string{headers["paypal-auth-algo"]},
		},
		Body: io.NopCloser(bytes.NewReader(body)),
	}

	resp, err := a.client.VerifyWebhookSignature(
		context.Background(),
		req,
		config.PayPalWebhookID,
	)
	if err != nil {
		return false, err
	}
	return resp.VerificationStatus == "SUCCESS", nil
}
