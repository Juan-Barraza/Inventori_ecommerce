package repository

import (
	"errors"
	domain "inventory/internal/fiber/domain/entities"
	"inventory/pkg"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *pkg.Database
}

func NewProductRespository(db *pkg.Database) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) AddProduct(product *domain.Product) error {
	return r.db.DB.Create(product).Error
}

func (r *ProductRepository) GetAllProducts(category string, providerId uint) (*gorm.DB, []domain.Product, error) {
	var products []domain.Product
	query := r.db.DB.Model(&domain.Product{}).Preload("Category").
		Preload("Provider").
		Preload("Images").
		Order("id asc")
	if category != "" {
		query = query.Joins("INNER JOIN categories ON products.category_id = categories.id").
			Where("categories.category_name ILIKE ?", "%"+category+"%")
	}

	if providerId > 0 {
		query = query.Where("provider_id = ?", providerId)
	}

	result := query.Find(&products)
	if result.Error != nil {
		return nil, nil, result.Error
	}

	return query, products, nil
}

func (r *ProductRepository) GetById(id uint) (*domain.Product, error) {
	var product domain.Product
	result := r.db.DB.Model(&domain.Product{}).Preload("Provider").Preload("Category").First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (r *ProductRepository) ValidateUniqueProduct(prod *domain.Product) error {
	var count int64

	err := r.db.DB.Model(&domain.Product{}).
		Where("name ILIKE ? AND category_id = ? AND provider_id = ?", "%"+prod.Name+"%", prod.CategoryId, prod.ProviderId).
		Count(&count).Error
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("already exist this product in any category and provider")
	}

	return nil
}

func (r *ProductRepository) DeleteProduct(product *domain.Product) error {
	return r.db.DB.Unscoped().Delete(product).Error
}

func (r *ProductRepository) UpdateProduct(product *domain.Product) error {
	return r.db.DB.Model(&product).Where("id = ?", product.ID).Updates(product).Error
}
