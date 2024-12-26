package provider

import (
	"fmt"
	domain "inventory/internal/fiber/domain/entities"
	"inventory/internal/fiber/domain/repositories"
)

type GetProviderService struct {
	proviRepo repositories.IProviderRepository
}

func NewGetProviderService(proviRepo repositories.IProviderRepository) *GetProviderService {
	return &GetProviderService{
		proviRepo: proviRepo,
	}
}

func (s *GetProviderService) GetALL() ([]domain.Provider, error) {
	provider, err := s.proviRepo.GetAllProvider()
	if err != nil {
		return nil, fmt.Errorf("error to get provider")
	}

	if provider == nil {
		provider = []domain.Provider{}
	}

	return provider, nil
}
