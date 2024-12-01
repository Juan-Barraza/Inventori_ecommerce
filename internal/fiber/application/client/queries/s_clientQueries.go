package queries

import (
	domain "inventory/internal/fiber/domain/entities"
	"inventory/internal/fiber/domain/repositories"
)

type ClientQuerysService struct {
	clientRepo repositories.IClientRepository
}

func NewClientQuerysService(clientRepo repositories.IClientRepository) *ClientQuerysService {
	return &ClientQuerysService{
		clientRepo: clientRepo,
	}
}

func (s *ClientQuerysService) GetAll() ([]domain.Client, error) {
	clients, err := s.clientRepo.GetAll()
	if err != nil {
		return nil, err
	}
	if clients == nil {
		clients = []domain.Client{}
	}

	return clients, nil
}
