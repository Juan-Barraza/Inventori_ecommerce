package routes

import (
	paymentfactory "inventory/internal/fiber/application/paymentFactory"
	"inventory/internal/fiber/application/transaction/commands"
	"inventory/internal/fiber/infrastructure/http/handlers"
	"inventory/internal/fiber/infrastructure/paidMethod/paypal"
	"inventory/internal/fiber/infrastructure/repository"
	"inventory/pkg"
	"log"

	"github.com/gofiber/fiber/v3"
)

func SetupTransactionRoutes(apiV1 fiber.Router, db *pkg.Database) {
	// Repositories
	txRepo := repository.NewTransactionRepository(db)
	gateway, err := paypal.NewPayPalAdapter()
	if err != nil {
		log.Fatal("failed to init PayPal adapter: %w", err)
	}
	factory := paymentfactory.NewPaymentFactory()
	createTransactionService := commands.NewCreateTransactionCommand(txRepo, factory)
	captureOrderPaypal := commands.NewCaptureOrderPaypal(txRepo, gateway)
	webHookService := commands.NewWebHookService(txRepo, gateway)
	handler := handlers.NewTransactionHandler(createTransactionService, webHookService, captureOrderPaypal)

	apiV1.Post("/orders/:orderId/transactions", handler.CreateTransaction)
	apiV1.Get("orders/:orderId/captureOrderPaypal", handler.CaptureOrderPaypal)
	// apiV1.Post("/transactions/webhook", handler.HandleWebHook)

	apiV1.Get("/payment/success", func(c fiber.Ctx) error {
		return c.SendString("Pago completado con éxito.")
	})
	apiV1.Get("/payment/cancel", func(c fiber.Ctx) error {
		return c.SendString("El pago fue cancelado.")
	})

}
