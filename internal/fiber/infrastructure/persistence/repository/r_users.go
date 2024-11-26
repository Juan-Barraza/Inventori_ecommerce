package repository

import (
	"inventory/internal/fiber/domain/models"
	"inventory/internal/fiber/infrastructure/persistence/mappers"
	modelsgorm "inventory/internal/fiber/infrastructure/persistence/modelsGORM"
	"inventory/pkg"
)

type UserRepository struct {
	db *pkg.Database
}

func NewUserRepository(db *pkg.Database) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(us *domain.User) error {
	user := mappers.ToUserGorm(us)
	return r.db.DB.Create(user).Error
}

func (r *UserRepository) GetAll() ([]domain.User, error) {
	var usersGorm []modelsgorm.User
	var users []domain.User
	err := r.db.DB.Model(modelsgorm.User{}).Find(&usersGorm).Error
	if err != nil {
		return nil, err
	}

	for _, user := range usersGorm {
		users = append(users, mappers.FromUserGorm(&user))
	}

	return users, nil
}
