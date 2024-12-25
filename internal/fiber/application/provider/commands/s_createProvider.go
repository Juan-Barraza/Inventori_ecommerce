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

type CreateProviderService struct {
	provRep repositories.IProviderRepository
	userRep repositories.IUserRepository
}

func NewCreateProviderService(provRep repositories.IProviderRepository,
	userRep repositories.IUserRepository) *CreateProviderService {
	return &CreateProviderService{
		provRep: provRep,
		userRep: userRep,
	}
}

func (s *CreateProviderService) CreateProvider(providerData *domain.Provider) error {
	providerExisting, err := s.provRep.GetByPhoneNummber(providerData.PhoneNumber)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if providerExisting != nil {
		return fmt.Errorf("provider with this phone number already exist")
	}

	userExist, err := s.userRep.FindByEmail(providerData.User.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if userExist != nil {
		return fmt.Errorf("email already exists")
	}

	password, err := utils.HashPassword(providerData.User.Password)
	if err != nil {
		log.Println("no se encrypto")
		return fmt.Errorf("error to encrytp password")
	}
	user := &domain.User{
		Email:    providerData.User.Email,
		Password: password,
	}
	err = s.userRep.Create(user)
	if err != nil {
		log.Println("no se creo el user")
		return fmt.Errorf("can not be create user")
	}

	providerData.UserID = user.ID
	if providerData.UserID == 0 {
		println("The user not asignement")
	}
	err = s.provRep.CreateProvider(providerData)
	if err != nil {
		return fmt.Errorf("can not be create provider")
	}

	return nil
}
