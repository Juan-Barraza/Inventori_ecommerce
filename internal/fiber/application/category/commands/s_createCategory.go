package commands

import (
	"fmt"
	domain "inventory/internal/fiber/domain/entities"
	"inventory/internal/fiber/domain/repositories"
)

type CreateCategoryService struct {
	cateRepo repositories.ICategory
}

func NewCreateCategoryService(cateRepo repositories.ICategory) *CreateCategoryService {
	return &CreateCategoryService{cateRepo: cateRepo}
}

func (s *CreateCategoryService) CreateCategory(category *domain.Category) error {
	err := s.cateRepo.ValidateCategory(category.CategoryName)
	if err != nil {
		return err
	}
	if err := s.cateRepo.AddCategory(category); err != nil {
		return fmt.Errorf("error to create category")
	}

	return nil
}
