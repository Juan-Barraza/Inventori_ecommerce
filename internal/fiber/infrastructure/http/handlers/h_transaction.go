package handlers

import (
	"context"
	"inventory/internal/fiber/application/transaction/commands"
	"inventory/internal/fiber/domain/ports"
	"inventory/pkg/utils/validators"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
)

type TransactionHandler struct {
	txService          *commands.CreateTransactionCommand
	webHookService     *commands.WebHookService
	captureOrderPaypal *commands.CaptureOrderPaypal
}

func NewTransactionHandler(txService *commands.CreateTransactionCommand,
	webHookService *commands.WebHookService,
	captureOrderPaypal *commands.CaptureOrderPaypal) *TransactionHandler {
	return &TransactionHandler{
		txService:          txService,
		webHookService:     webHookService,
		captureOrderPaypal: captureOrderPaypal,
	}
}

func (h *TransactionHandler) CreateTransaction(c fiber.Ctx) error {
	var req ports.PaymentRequest
	orderIdStr := c.Params("orderId")
	orderID, err := strconv.Atoi(orderIdStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "error to convert string to int"})
	}

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if !validators.IsValidPaymentMethod(req.Method) {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid payment method"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	status, url, err := h.txService.ProcessPaymentTransaction(ctx, uint(orderID), req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"status": status, "url": url})
}

func (h *TransactionHandler) CaptureOrderPaypal(c fiber.Ctx) error {
	orderIdSTr := c.Params("orderId")
	orderId, err := strconv.Atoi(orderIdSTr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Error to convert string to int"})
	}

	status, err := h.captureOrderPaypal.CaptureOrderPaypal(uint(orderId))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"Status": status})
}

// func (h *TransactionHandler) HandleWebHook(c fiber.Ctx) error {
// 	body := c.Body()
// 	headers := make(map[string]string)
// 	c.Request().Header.VisitAll(func(key, value []byte) {
// 		headers[string(key)] = string(value)
// 	})

// 	responseMessaage, err := h.webHookService.HandleWebHook(context.Background(), body, headers)
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	return c.Status(200).JSON(fiber.Map{"message": responseMessaage})
// }
