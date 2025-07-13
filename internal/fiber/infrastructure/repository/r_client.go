package repository

import (
	domain "inventory/internal/fiber/domain/entities"
	"inventory/pkg"

	"gorm.io/gorm"
)

type ClientRepository struct {
	db *pkg.Database
}

func NewClientRepository(db *pkg.Database) *ClientRepository {
	return &ClientRepository{db: db}
}

func (r *ClientRepository) Create(client *domain.Client) error {
	return r.db.DB.Create(client).Error
}

func (r *ClientRepository) GetAll() (*gorm.DB, error) {
	query := r.db.DB.Model(&domain.Client{}).Preload("User")

	return query, nil
}

func (r *ClientRepository) GetById(id uint) (*domain.Client, error) {
	var client *domain.Client

	err := r.db.DB.Model(&domain.Client{}).Preload("User").First(&client, id).Error
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (r *ClientRepository) Update(client *domain.Client) error {
	return r.db.DB.Model(&client).Where("id = ?", client.ID).Updates(client).Error
}

func (r *ClientRepository) Delete(client *domain.Client) error {
	return r.db.DB.Unscoped().Delete(client).Error
}

func (r *ClientRepository) FindByDocumentNumber(documentNumber string) (*domain.Client, error) {
	var client *domain.Client
	err := r.db.DB.Model(&domain.Client{}).Where("document_number = ?", documentNumber).First(&client).Error
	if err != nil {
		return nil, err
	}
	return client, nil
}
