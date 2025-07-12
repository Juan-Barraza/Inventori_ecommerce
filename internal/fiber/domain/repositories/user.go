package repositories

import (
	domain "inventory/internal/fiber/domain/entities"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user *domain.User) error
	GetAll() (*gorm.DB, error)
	Update(user *domain.User) error
	Delete(user *domain.User) error
	FindByID(id uint) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
}
