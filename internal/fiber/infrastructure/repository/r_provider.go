package repository

import (
	domain "inventory/internal/fiber/domain/entities"
	"inventory/pkg"

	"gorm.io/gorm"
)

type ProviderRepository struct {
	db *pkg.Database
}

func NewProviderRepsoitor(db *pkg.Database) *ProviderRepository {
	return &ProviderRepository{db: db}
}

func (r *ProviderRepository) CreateProvider(provider *domain.Provider) error {
	return r.db.DB.Create(provider).Error
}

func (r *ProviderRepository) GetAllProvider() (*gorm.DB, error) {
	query := r.db.DB.Model(&domain.Provider{}).Preload("User")

	return query, nil
}

func (r *ProviderRepository) GetById(id uint) (*domain.Provider, error) {
	var provider *domain.Provider

	err := r.db.DB.Model(domain.Provider{}).Preload("User").First(&provider, id).Error
	if err != nil {
		return nil, err
	}

	return provider, nil

}

func (r *ProviderRepository) Update(provider *domain.Provider) error {
	return r.db.DB.Model(&provider).Where("id = ?", provider.ID).Updates(provider).Error
}

func (r *ProviderRepository) Delete(provider *domain.Provider) error {
	return r.db.DB.Unscoped().Delete(provider).Error
}

func (r *ProviderRepository) GetByPhoneNummber(number string) (*domain.Provider, error) {
	var provider *domain.Provider
	err := r.db.DB.Model(&domain.Provider{}).Where("phone_number = ?", number).First(&provider).Error
	if err != nil {
		return nil, err
	}

	return provider, nil
}
