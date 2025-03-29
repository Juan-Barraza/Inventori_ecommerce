package repository

import (
	"errors"
	domain "inventory/internal/fiber/domain/entities"
	"inventory/pkg"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *pkg.Database
}

func NewCategoryRepository(db *pkg.Database) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) AddCategory(category *domain.Category) error {
	return r.db.DB.Create(category).Error
}

func (r *CategoryRepository) GetAllCategories() (*gorm.DB, error) {
	query := r.db.DB.Model(&domain.Category{})
	return query, nil
}

func (r *CategoryRepository) GetById(id uint) (*domain.Category, error) {
	var category *domain.Category
	if err := r.db.DB.Model(&domain.Category{}).
		First(&category, id).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (r *CategoryRepository) DeleteCategory(category *domain.Category) error {
	return r.db.DB.Unscoped().Delete(category).Error
}

func (r *CategoryRepository) ValidateCategory(name string) error {
	var count int64
	err := r.db.DB.Model(&domain.Category{}).
		Where("category_name ILIKE ?", "%"+name+"%").
		Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("already exist this category")
	}

	return nil
}
