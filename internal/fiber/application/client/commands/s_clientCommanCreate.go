package commands

import (
	"errors"
	"fmt"
	domain "inventory/internal/fiber/domain/entities"
	"inventory/internal/fiber/domain/repositories"
	"inventory/pkg/utils"
	"log"

	"gorm.io/gorm"
)

type CreateClientCommandsService struct {
	clientRepo repositories.IClientRepository
	userRepo   repositories.IUserRepository
}

func NewClientCommandsService(clientRepo repositories.IClientRepository,
	userRepo repositories.IUserRepository) *CreateClientCommandsService {
	return &CreateClientCommandsService{
		clientRepo: clientRepo,
		userRepo:   userRepo,
	}
}

func (s *CreateClientCommandsService) CreateClient(client *domain.Client) error {
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
	client.User = domain.User{}
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
