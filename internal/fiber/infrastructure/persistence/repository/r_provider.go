package repository

import (
	domain "inventory/internal/fiber/domain/entities"
	"inventory/internal/fiber/infrastructure/persistence/mappers"
	modelsgorm "inventory/internal/fiber/infrastructure/persistence/modelsGORM"
	"inventory/pkg"

	"gorm.io/gorm"
)

type ProviderRepository struct {
	db *pkg.Database
}

func NewProviderRepsoitor(db *pkg.Database) *ProviderRepository {
	return &ProviderRepository{db: db}
}

func (r *ProviderRepository) CreateProvider(prov *domain.Provider) error {
	var client = mappers.ToProviderGorm(prov)
	return r.db.DB.Create(client).Error
}

func (r *ProviderRepository) GetAllProvider() (*gorm.DB, error) {
	query := r.db.DB.Model(&modelsgorm.Provider{}).Preload("User")

	return query, nil
}

func (r *ProviderRepository) GetById(id uint) (*domain.Provider, error) {
	var providerGorm modelsgorm.Provider

	err := r.db.DB.Model(modelsgorm.Provider{}).Preload("User").First(&providerGorm, id).Error
	if err != nil {
		return nil, err
	}

	provider := mappers.FromProviderGorm(&providerGorm)

	return provider, nil

}

func (r *ProviderRepository) Update(prov *domain.Provider) error {
	provider := mappers.ToProviderGorm(prov)
	provider.ID = prov.ID
	return r.db.DB.Model(&provider).Where("id = ?", provider.ID).Updates(provider).Error
}

func (r *ProviderRepository) Delete(prov *domain.Provider) error {
	provider := mappers.ToProviderGorm(prov)
	provider.ID = prov.ID
	return r.db.DB.Unscoped().Delete(provider).Error
}

func (r *ProviderRepository) GetByPhoneNummber(number string) (*domain.Provider, error) {
	var providerGorm modelsgorm.Provider
	err := r.db.DB.Model(&modelsgorm.Provider{}).Where("phone_number = ?", number).First(&providerGorm).Error
	if err != nil {
		return nil, err
	}
	provider := mappers.FromProviderGorm(&providerGorm)

	return provider, nil
}
