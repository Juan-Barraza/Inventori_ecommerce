package commands

import (
	"fmt"
	domain "inventory/internal/fiber/domain/entities"
	"inventory/internal/fiber/domain/repositories"
	"inventory/pkg/utils"
)


type UpdateClientCommandsService struct {
	clientRepo repositories.IClientRepository
	userRepo   repositories.IUserRepository
}

func NewUpdateClientCommandsService(clientRepo repositories.IClientRepository,
	userRepo repositories.IUserRepository) *UpdateClientCommandsService {
	return &UpdateClientCommandsService{
		clientRepo: clientRepo,
		userRepo:   userRepo,
	}
}

func (s *UpdateClientCommandsService) UpdateClient(id uint, clienData *domain.Client) error {
	existingClient, err := s.clientRepo.GetById(id)
	if err != nil {
		return fmt.Errorf("client not found")
	}

	if clienData.Name != "" {
		existingClient.Name = clienData.Name
	}

	if clienData.LastName != "" {
		existingClient.LastName = clienData.LastName
	}

	if clienData.DocumentNumber != "" {
		existingClient.DocumentNumber = clienData.DocumentNumber
	}

	if clienData.Address != "" {
		existingClient.Address = clienData.Address
	}

	if clienData.TypeDocument != "" {
		existingClient.TypeDocument = clienData.TypeDocument
	}

	if clienData.PhoneNumber != "" {
		existingClient.PhoneNumber = clienData.PhoneNumber
	}

	existUser, err := s.userRepo.FindByID(existingClient.UserID)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	if clienData.User.Email != "" {
		existUser.Email = clienData.User.Email
	}

	if clienData.User.Password != "" {
		paswd, err := utils.HashPassword(clienData.User.Password)
		if err != nil {
			return err
		}

		existUser.Password = paswd
	}

	if err = s.userRepo.Update(existUser); err != nil {
		return fmt.Errorf("error to update user")
	}

	if err = s.clientRepo.Update(existingClient); err != nil {
		return fmt.Errorf("error to update client")
	}
	return nil
}
