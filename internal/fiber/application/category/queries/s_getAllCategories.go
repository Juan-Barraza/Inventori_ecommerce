package queries

import (
	"fmt"
	domain "inventory/internal/fiber/domain/entities"
	"inventory/internal/fiber/domain/repositories"
	"inventory/internal/fiber/infrastructure/repository"
	"inventory/pkg/utils"
)

type GetAllCategoriesService struct {
	categoryRepo  repositories.ICategory
	paginationRep *repository.PaginationRepository
}

func NewGetAllCategoriesService(categoryRepo repositories.ICategory,
	paginationRep *repository.PaginationRepository) *GetAllCategoriesService {
	return &GetAllCategoriesService{categoryRepo: categoryRepo}
}

func (s *GetAllCategoriesService) GetAllCategories(pagination *utils.Pagination) (*utils.Pagination, error) {
	query, err := s.categoryRepo.GetAllCategories()
	if err != nil {
		return nil, fmt.Errorf("error to get product")
	}

	var categories []domain.Category
	var categoriesJson []domain.CategoryJson
	paginationResult, err := s.paginationRep.GetPaginatedResults(query, pagination, &categories)
	if err != nil {
		return nil, fmt.Errorf("error to make pagination")
	}

	for _, category := range categories {
		categoriesJson = append(categoriesJson, *domain.ToCategory(&category))
	}

	paginationResult.Data = categoriesJson

	return paginationResult, nil
}
