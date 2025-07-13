package application

import (
	"fmt"
	domain "inventory/internal/fiber/domain/entities"
	"inventory/internal/fiber/domain/repositories"
	"inventory/internal/fiber/infrastructure/repository"
	"inventory/pkg/utils"
)

type UserService struct {
	userRepository repositories.IUserRepository
	paginationRep  *repository.PaginationRepository
}

func NewUserService(userRepo repositories.IUserRepository,
	paginationRep *repository.PaginationRepository) *UserService {
	return &UserService{userRepository: userRepo, paginationRep: paginationRep}
}

func (s *UserService) RegisterUser(usersData domain.User) error {

	return s.userRepository.Create(&usersData)
}

func (s *UserService) GelAll(pagination *utils.Pagination) (*utils.Pagination, error) {
	query, err := s.userRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf("error to get users")
	}
	var users []domain.User
	var usersResponse []domain.UserResponse
	paginationResult, err := s.paginationRep.GetPaginatedResults(query, pagination, &users)
	if err != nil {
		return nil, fmt.Errorf("error to get pagination")
	}
	for _, user := range users {
		usersResponse = append(usersResponse, *domain.ToUser(&user))
	}
	paginationResult.Data = usersResponse

	return paginationResult, nil
}
