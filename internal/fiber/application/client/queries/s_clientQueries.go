package queries

import (
	"fmt"
	domain "inventory/internal/fiber/domain/entities"
	"inventory/internal/fiber/domain/repositories"
	"inventory/internal/fiber/infrastructure/repository"
	"inventory/pkg/utils"
)

type ClientQuerysService struct {
	clientRepo    repositories.IClientRepository
	paginationRep *repository.PaginationRepository
}

func NewClientQuerysService(clientRepo repositories.IClientRepository,
	paginationRep *repository.PaginationRepository,
) *ClientQuerysService {
	return &ClientQuerysService{
		clientRepo:    clientRepo,
		paginationRep: paginationRep,
	}
}

func (s *ClientQuerysService) GetAll(pagination *utils.Pagination) (*utils.Pagination, error) {
	query, err := s.clientRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("error to get client")
	}
	var clients []domain.Client
	var clientsJson []domain.ClientResponse
	paginationResult, err := s.paginationRep.GetPaginatedResults(query, pagination, &clients)
	if err != nil {
		return nil, fmt.Errorf("error to get pagintaion")
	}

	for _, client := range clients {
		clientsJson = append(clientsJson, *domain.ToClient(&client))
	}

	paginationResult.Data = clientsJson

	return paginationResult, nil
}
