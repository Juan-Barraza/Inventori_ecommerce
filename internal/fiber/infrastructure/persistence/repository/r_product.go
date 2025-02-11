package repository

import (
	domain "inventory/internal/fiber/domain/entities"
	"inventory/internal/fiber/infrastructure/persistence/mappers"
	modelsgorm "inventory/internal/fiber/infrastructure/persistence/modelsGORM"
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
	productGorm := mappers.ToProductGorm(product)
	return r.db.DB.Create(productGorm).Error
}

func (r *ProductRepository) GetAllProducts(category string, providerId uint) (*gorm.DB, []domain.Product, error) {
	var productsGorm []modelsgorm.Product
	var products []domain.Product
	query := r.db.DB.Model(&modelsgorm.Product{}).Preload("Category").
		Preload("Provider").
		Order("id asc")
	if category != "" {
		query = query.Joins("INNER JOIN categories ON products.category_id = categories.id").Where("categories.CategoryName ILIKE= ?", "%"+category+"%")
	}

	if providerId > 0 {
		query = query.Where("provider_id = ?", providerId)
	}

	if err := query.Find(&productsGorm); err != nil {
		return nil, nil, err.Error
	}

	for _, product := range productsGorm {
		products = append(products, *mappers.FromProductGorm(&product))
	}

	return query, products, nil
}

func (r *ProductRepository) GetById(id uint) (*domain.Product, error) {
	var producGorm modelsgorm.Product
	if err := r.db.DB.Model(&modelsgorm.Product{}).Preload("Provider").Preload("Category").First(&producGorm); err != nil {
		return nil, err.Error
	}
	product := mappers.FromProductGorm(&producGorm)

	return product, nil
}

func (r *ProductRepository) DeleteProduc(prod *domain.Product) error {
	product := mappers.ToProductGorm(prod)
	product.ID = prod.ID
	return r.db.DB.Unscoped().Delete(product).Error
}
