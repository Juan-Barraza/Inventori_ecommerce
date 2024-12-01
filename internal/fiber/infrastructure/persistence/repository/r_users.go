package repository

import (
	domain "inventory/internal/fiber/domain/entities"
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
	if err := r.db.DB.Create(user).Error; err != nil {
		return err
	}
	us.ID = user.ID
	return nil
}

func (r *UserRepository) GetAll() ([]domain.User, error) {
	var usersGorm []modelsgorm.User
	var users []domain.User
	err := r.db.DB.Model(modelsgorm.User{}).Find(&usersGorm).Error
	if err != nil {
		return nil, err
	}

	for _, user := range usersGorm {
		users = append(users, *mappers.FromUserGorm(&user))
	}

	return users, nil
}

func (r *UserRepository) Update(usr *domain.User) error {
	user := mappers.ToUserGorm(usr)
	user.ID = usr.ID
	return r.db.DB.Save(&user).Error
}

func (r *UserRepository) Delete(us *domain.User) error {
	user := mappers.ToUserGorm(us)
	user.ID = us.ID
	return r.db.DB.Unscoped().Delete(user).Error
}

func (r *UserRepository) FindByID(id uint) (*domain.User, error) {
	var userGorm *modelsgorm.User

	if err := r.db.DB.Model(modelsgorm.User{}).First(&userGorm, id).Error; err != nil {
		return nil, err
	}

	user := mappers.FromUserGorm(userGorm)

	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	var userGorm modelsgorm.User
	err := r.db.DB.Model(modelsgorm.User{}).Where("email = ?", email).First(&userGorm).Error
	if err != nil {
		return nil, err
	}

	user := mappers.FromUserGorm(&userGorm)

	return user, nil
}
