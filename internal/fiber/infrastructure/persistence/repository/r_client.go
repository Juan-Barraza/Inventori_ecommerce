package repository

import (
	domain "inventory/internal/fiber/domain/entities"
	"inventory/internal/fiber/infrastructure/persistence/mappers"
	modelsgorm "inventory/internal/fiber/infrastructure/persistence/modelsGORM"
	"inventory/pkg"
)

type ClientRepository struct {
	db *pkg.Database
}

func NewClientRepository(db *pkg.Database) *ClientRepository {
	return &ClientRepository{db: db}
}

func (r *ClientRepository) Create(cli *domain.Client) error {
	var client = mappers.ToClientGorm(cli)
	return r.db.DB.Create(client).Error
}

func (r *ClientRepository) GetAll() ([]domain.Client, error) {
	var clientsGorm []modelsgorm.Client
	var clients []domain.Client
	err := r.db.DB.Preload("User").Find(&clientsGorm).Error
	if err != nil {
		return nil, err
	}
	for _, client := range clientsGorm {
		clients = append(clients, *mappers.FromClientGorm(&client))
	}

	return clients, nil
}

func (r *ClientRepository) GetById(id uint) (*domain.Client, error) {
	var clientGorm modelsgorm.Client

	err := r.db.DB.Model(modelsgorm.Client{}).Preload("User").First(&clientGorm, id).Error
	if err != nil {
		return nil, err
	}

	client := mappers.FromClientGorm(&clientGorm)

	return client, nil

}

func (r *ClientRepository) Update(cli *domain.Client) error {
	client := mappers.ToClientGorm(cli)
	client.ID = cli.ID
	return r.db.DB.Model(&client).Where("id = ?", client.ID).Updates(client).Error
}

func (r *ClientRepository) Delete(cli *domain.Client) error {
	client := mappers.ToClientGorm(cli)
	client.ID = cli.ID
	return r.db.DB.Unscoped().Delete(client).Error
}

func (r *ClientRepository) FindByDocumentNumber(documentNumber string) (*domain.Client, error) {
	var clientGorm modelsgorm.Client
	err := r.db.DB.Model(&modelsgorm.Client{}).Where("document_number = ?", documentNumber).First(&clientGorm).Error
	if err != nil {
		return nil, err
	}
	client := mappers.FromClientGorm(&clientGorm)

	return client, nil
}
