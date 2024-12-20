package mappers

import (
	domain "inventory/internal/fiber/domain/entities"
	modelsgorm "inventory/internal/fiber/infrastructure/persistence/modelsGORM"
)

func ToUserGorm(u *domain.User) *modelsgorm.User {
	var user = &modelsgorm.User{
		Email:    u.Email,
		Password: u.Password,
	}
	return user
}

func FromUserGorm(ug *modelsgorm.User) *domain.User {
	return &domain.User{ID: ug.ID, Email: ug.Email, Password: ug.Password}
}
