package commands

import (
	"errors"
	"fmt"
	"inventory/internal/fiber/domain/repositories"
)

var (
	ErrProviderNotFound = errors.New("provider not found")
	ErrUserNotFound     = errors.New("user not found")
)

type DeleteProviderService struct {
	providerRep repositories.IProviderRepository
	userRepo    repositories.IUserRepository
}

func NewDeleteProviderService(providerRep repositories.IProviderRepository,
	userRepo repositories.IUserRepository) *DeleteProviderService {
	return &DeleteProviderService{
		providerRep: providerRep,
		userRepo:    userRepo,
	}
}

func (s *DeleteProviderService) DeleteProvider(id uint) error {
	provider, err := s.providerRep.GetById(id)
	if err != nil {
		return ErrProviderNotFound
	}

	user, err := s.userRepo.FindByID(provider.UserID)
	if err != nil {
		return ErrUserNotFound
	}

	if err = s.providerRep.Delete(provider); err != nil {
		return fmt.Errorf("error deleting provider")
	}
	if err = s.userRepo.Delete(user); err != nil {
		return fmt.Errorf("error deleting user")
	}

	return nil
}
