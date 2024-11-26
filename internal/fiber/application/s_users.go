package application

import (
	domain "inventory/internal/fiber/domain/models"
	"inventory/internal/fiber/domain/repositories"
)

type UserService struct {
	userRepository repositories.IUserRepository
}

func NewUserService(userRepo repositories.IUserRepository) *UserService {
	return &UserService{userRepository: userRepo}
}

func (s *UserService) RegisterUser(usersData domain.User) error {

	return s.userRepository.Create(&usersData)
}

func (s *UserService) GelAll() ([]domain.User, error) {
	return s.userRepository.GetAll()
}
