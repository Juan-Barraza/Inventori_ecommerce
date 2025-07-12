package provider

import (
	"fmt"
	domain "inventory/internal/fiber/domain/entities"
	"inventory/internal/fiber/domain/repositories"
	"inventory/internal/fiber/infrastructure/repository"
	"inventory/pkg/utils"
)

type GetProviderService struct {
	proviRepo     repositories.IProviderRepository
	paginationRep *repository.PaginationRepository
}

func NewGetProviderService(proviRepo repositories.IProviderRepository,
	paginationRep *repository.PaginationRepository) *GetProviderService {
	return &GetProviderService{
		proviRepo:     proviRepo,
		paginationRep: paginationRep,
	}
}

func (s *GetProviderService) GetALL(pagination *utils.Pagination) (*utils.Pagination, error) {
	query, err := s.proviRepo.GetAllProvider()
	if err != nil {
		return nil, fmt.Errorf("error to get provider")
	}
	var providers []domain.Provider
	var providerJson []domain.ProviderResponse
	paginationResult, err := s.paginationRep.GetPaginatedResults(query, pagination, &providers)
	if err != nil {
		return nil, fmt.Errorf("error to get pagintaion")
	}
	for _, provider := range providers {
		providerJson = append(providerJson, *domain.ToProvider(&provider))
	}

	paginationResult.Data = providerJson

	return paginationResult, nil
}
