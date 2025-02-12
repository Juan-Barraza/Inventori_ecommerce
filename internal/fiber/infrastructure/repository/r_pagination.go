package repository

import (
	"inventory/pkg"
	"inventory/pkg/utils"

	"gorm.io/gorm"
)

type PaginationRepository struct {
	db *pkg.Database
}

func NewPaginationRepository(db *pkg.Database) *PaginationRepository {
	return &PaginationRepository{db: db}
}

func (r *PaginationRepository) GetPaginatedResults(query *gorm.DB, pagination *utils.Pagination, result interface{}) (*utils.Pagination, error) {
	var totalItems int64

	if err := query.Count(&totalItems).Error; err != nil {
		return nil, err
	}

	pagination.TotalItems = int(totalItems)

	offset := (pagination.Page - 1) * pagination.PageSize

	if err := query.Limit(pagination.PageSize).Offset(offset).Find(result).Error; err != nil {
		return nil, err
	}

	pagination.Calculate()

	pagination.Data = result

	return pagination, nil
}
