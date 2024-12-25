package mappers

import (
	domain "inventory/internal/fiber/domain/entities"
	modelsgorm "inventory/internal/fiber/infrastructure/persistence/modelsGORM"
)

func ToProviderGorm(p *domain.Provider) *modelsgorm.Provider {
	return &modelsgorm.Provider{
		Name:          p.Name,
		Address:       p.Address,
		PhoneNumber:   p.PhoneNumber,
		TypeOfProduct: p.TypeOfProduct,
		UserID:        p.UserID,
	}
}

func FromProviderGorm(p *modelsgorm.Provider) *domain.Provider {
	return &domain.Provider{
		Name:          p.Name,
		Address:       p.Address,
		PhoneNumber:   p.PhoneNumber,
		TypeOfProduct: p.TypeOfProduct,
		UserID:        p.UserID,
		User: domain.User{
			Email:    p.User.Email,
			Password: p.User.Password,
		},
	}
}
