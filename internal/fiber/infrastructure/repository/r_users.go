package repository

import (
	domain "inventory/internal/fiber/domain/entities"
	"inventory/pkg"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *pkg.Database
}

func NewUserRepository(db *pkg.Database) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *domain.User) error {
	if err := r.db.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetAll() (*gorm.DB, error) {
	query := r.db.DB.Model(&domain.User{})

	return query, nil
}

func (r *UserRepository) Update(user *domain.User) error {
	return r.db.DB.Save(&user).Error
}

func (r *UserRepository) Delete(user *domain.User) error {
	return r.db.DB.Unscoped().Delete(user).Error
}

func (r *UserRepository) FindByID(id uint) (*domain.User, error) {
	var user *domain.User

	if err := r.db.DB.Model(domain.User{}).First(&user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	var user *domain.User
	err := r.db.DB.Model(domain.User{}).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
