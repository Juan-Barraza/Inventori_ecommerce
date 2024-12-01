package commands

import (
	"fmt"
	"inventory/internal/fiber/domain/repositories"
)

type DeleteClientCommandsService struct {
	clientRepo repositories.IClientRepository
	userRepo   repositories.IUserRepository
}

func NewDeleteClientCommandsService(clientRepo repositories.IClientRepository,
	userRepo repositories.IUserRepository) *DeleteClientCommandsService {
	return &DeleteClientCommandsService{
		clientRepo: clientRepo,
		userRepo:   userRepo,
	}
}


func (s *DeleteClientCommandsService) DeleteClient(id uint) error {
	client, err := s.clientRepo.GetById(id)
	if err != nil {
		return fmt.Errorf("error to finding client")
	}

	user, err := s.userRepo.FindByID(client.UserID)
	if err != nil {
		return fmt.Errorf("error to finding user")
	}
	err = s.userRepo.Delete(user)
	if err != nil {
		return fmt.Errorf("error deleting user")
	}
	err = s.clientRepo.Delete(client)
	if err != nil {
		return fmt.Errorf("error deleting client")
	}

	return nil
}
