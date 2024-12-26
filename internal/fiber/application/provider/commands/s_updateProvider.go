package commands

import (
	"fmt"
	domain "inventory/internal/fiber/domain/entities"
	"inventory/internal/fiber/domain/repositories"
	"inventory/pkg/utils"
)

type UpdateProviderService struct {
	providerRepo repositories.IProviderRepository
	userRepo     repositories.IUserRepository
}

func NewUpdateProviderService(providerRepo repositories.IProviderRepository,
	userRepo repositories.IUserRepository) *UpdateProviderService {
	return &UpdateProviderService{
		providerRepo: providerRepo,
		userRepo:     userRepo,
	}
}

func (s *UpdateProviderService) UpdateProvider(id uint, providerData *domain.Provider) error {
	existingProvide, err := s.providerRepo.GetById(id)
	if err != nil {
		return fmt.Errorf("provider do not exist")
	}

	if providerData.Name != "" {
		existingProvide.Name = providerData.Name
	}
	if providerData.Address != "" {
		existingProvide.Address = providerData.Address
	}
	if providerData.PhoneNumber != "" {
		existingProvide.PhoneNumber = providerData.PhoneNumber
	}
	if providerData.TypeOfProduct != "" {
		existingProvide.TypeOfProduct = providerData.TypeOfProduct
	}

	existingUser, err := s.userRepo.FindByID(existingProvide.UserID)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	if providerData.User.Email != "" {
		existingUser.Email = providerData.User.Email
	}
	if providerData.User.Password != "" {
		paswd, err := utils.HashPassword(providerData.User.Password)
		if err != nil {
			return fmt.Errorf("fail to encrytp password")
		}
		existingUser.Password = paswd
	}

	if err = s.userRepo.Update(existingUser); err != nil {
		return fmt.Errorf("error to update user")
	}
	if err = s.providerRepo.Update(existingProvide); err != nil {
		return fmt.Errorf("error to update provider")
	}

	return nil
}
