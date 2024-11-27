package application

import (
	"errors"
	"fmt"
	domain "inventory/internal/fiber/domain/models"
	"inventory/internal/fiber/domain/repositories"
	"inventory/pkg/utils"
	"log"

	"gorm.io/gorm"
)

type ClientService struct {
	clientRepo repositories.IClientRepository
	userRepo   repositories.IUserRepository
}

func NewClientService(clientRepo repositories.IClientRepository,
	userRepo repositories.IUserRepository) *ClientService {
	return &ClientService{
		clientRepo: clientRepo,
		userRepo:   userRepo,
	}
}

func (s *ClientService) CreateClient(client *domain.Client) error {
	userExist, err := s.userRepo.FindByEmail(client.User.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if userExist != nil {
		return fmt.Errorf("email already exists")
	}

	clientExist, err := s.clientRepo.FindByDocumentNumber(client.DocumentNumber)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {

		return err
	}
	if clientExist != nil {
		return fmt.Errorf("client Already exists")
	}
	password, err := utils.HashPassword(client.User.Password)
	if err != nil {
		log.Println("no se encrypto")
		return err
	}
	user := &domain.User{
		Email:    client.User.Email,
		Password: password,
	}
	err = s.userRepo.Create(user)
	if err != nil {
		log.Println("no se creo el user")
		return err
	}

	log.Println("si se creo el user")
	client.UserID = user.ID
	if client.UserID == 0 {
		println("the user not asignement")
	}

	err = s.clientRepo.Create(client)
	if err != nil {
		log.Println("no se creo el client")
		return err
	}

	return nil
}

func (s *ClientService) GetAll() ([]domain.Client, error) {
	return s.clientRepo.GetAll()
}
