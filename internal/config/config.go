package config

import (
	"inventory/pkg"
	"os"
	"time"
)

var (
	PayPalClientID     = os.Getenv("PAYPAL_CLIENT_ID")
	PayPalSecret       = os.Getenv("PAYPAL_SECRET")
	PayPalAPIBase      = os.Getenv("PAYPAL_API_BASE") // sandbox o live
	PayPalWebhookID    = os.Getenv("PAYPAL_WEBHOOK_ID")
	PayPalVerifyTicker = time.Minute * 5
)

func SetConfig() (*pkg.Database, error) {

	db, err := CreateTables()
	if err != nil {
		return nil, err
	}

	return db, nil
}