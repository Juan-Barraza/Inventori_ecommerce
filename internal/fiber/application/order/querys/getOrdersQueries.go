package querys

import (
	domain "inventory/internal/fiber/domain/entities"
	"inventory/internal/fiber/domain/repositories"
	"inventory/internal/fiber/infrastructure/repository"
	"inventory/pkg/utils"
)

type GetOrdersQuery struct {
	orderRepo     repositories.IOrderRepository
	paginationRep *repository.PaginationRepository
}

func NewGetOrdersQuery(orderRepo repositories.IOrderRepository,
	paginationRe *repository.PaginationRepository) *GetOrdersQuery {
	return &GetOrdersQuery{
		orderRepo:     orderRepo,
		paginationRep: paginationRe}
}

func (s *GetOrdersQuery) GetOrders(clientId, productId uint, pagination *utils.Pagination) (*utils.Pagination, error) {
	query, orders, err := s.orderRepo.GetOrders(clientId, productId)
	if err != nil {
		return nil, err
	}
	paginationResult, err := s.paginationRep.GetPaginatedResults(query, pagination, &orders)
	if err != nil {
		return nil, err
	}
	if orders == nil {
		orders = []domain.Order{}
	}
	var ordersJson []domain.OrderJson

	for _, order := range orders {
		ordersJson = append(ordersJson, *domain.ToOrder(&order))
	}

	paginationResult.Data = ordersJson
	return paginationResult, nil
}
