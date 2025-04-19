package routes

import (
	"inventory/internal/fiber/application/order/commands"
	"inventory/internal/fiber/application/order/querys"
	"inventory/internal/fiber/infrastructure/http/handlers"
	"inventory/internal/fiber/infrastructure/repository"
	"inventory/pkg"

	"github.com/gofiber/fiber/v3"
)

func SetOrderRoutes(apiV1 fiber.Router, db *pkg.Database) {
	paginationRep := repository.NewPaginationRepository(db)
	orderRepo := repository.NewOrderRepository(db)
	createOrderCommand := commands.NewCreateOrderCommand(orderRepo)
	getOrderQuery := querys.NewGetOrdersQuery(orderRepo, paginationRep)
	deleteOrderCommand := commands.NewDeleteOrderCommand(orderRepo)
	handlerOrder := handlers.NewOrderHandler(createOrderCommand, getOrderQuery, deleteOrderCommand)

	apiV1.Post("/orders", handlerOrder.CreateOrder)
	apiV1.Get("/orders", handlerOrder.GetOrders)
	apiV1.Delete("/orders/:id", handlerOrder.DeleteOrder)
}
