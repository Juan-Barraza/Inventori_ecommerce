package handlers

import (
	"inventory/internal/fiber/application/order/commands"
	"inventory/internal/fiber/application/order/querys"
	domain "inventory/internal/fiber/domain/entities"
	"inventory/pkg/utils"
	"inventory/pkg/utils/validators"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type OrderHandler struct {
	OrderCreateService *commands.CreateOrderCommand
	OrderGetService    *querys.GetOrdersQuery
	OrderDeleteService *commands.DeleteOrderCommand
}

func NewOrderHandler(orderService *commands.CreateOrderCommand,
	OrderGetService *querys.GetOrdersQuery,
	OrderDeleteService *commands.DeleteOrderCommand) *OrderHandler {
	return &OrderHandler{
		OrderCreateService: orderService,
		OrderGetService:    OrderGetService,
		OrderDeleteService: OrderDeleteService,
	}
}

func (h *OrderHandler) CreateOrder(c fiber.Ctx) error {
	var orderData *domain.OrderDTO
	if err := c.Bind().Body(&orderData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := validators.ValidateOrder(orderData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	products := make([]domain.Product, len(orderData.Products))
	for i, id := range orderData.Products {
		products[i] = domain.Product{Model: gorm.Model{ID: id}}

	}

	order := &domain.Order{
		Status:      orderData.Status,
		Quantity:    orderData.Quantity,
		Date:        orderData.Date,
		Description: orderData.Description,
		ClientId:    orderData.ClientId,
		Products:    products,
	}

	if err := h.OrderCreateService.CreateOrder(order); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create order"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Order created successfully"})
}

func (h *OrderHandler) GetOrders(c fiber.Ctx) error {
	pagination := c.Locals("pagination").(*utils.Pagination)
	clientIdStr := c.Query("clientId")
	productIdStr := c.Query("productId")
	var clientId, productId int
	var err error

	if clientIdStr != "" {
		clientId, err = strconv.Atoi(clientIdStr)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "error to parse id",
			})
		}

	}
	if productIdStr != "" {
		productId, err = strconv.Atoi(productIdStr)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "error to parse id",
			})
		}
	}

	paginatedOrders, err := h.OrderGetService.GetOrders(uint(clientId), uint(productId), pagination)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get orders"})
	}
	return c.JSON(paginatedOrders)
}

func (h *OrderHandler) DeleteOrder(c fiber.Ctx) error {
	orderId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id parameter",
		})
	}

	if err := h.OrderDeleteService.DeleteOrder(uint(orderId)); err != nil {
		if strings.Contains(err.Error(), "error checking order existence") {
			return c.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(204)

}
