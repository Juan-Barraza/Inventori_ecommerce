package commands

import (
	"fmt"
	"inventory/internal/fiber/domain/repositories"
)

type DeleteCategoryService struct {
	catRepo repositories.ICategory
}

func NewDeleteCategoryService(catRepo repositories.ICategory) *DeleteCategoryService {
	return &DeleteCategoryService{catRepo: catRepo}
}

func (s *DeleteCategoryService) DeleteCategory(id uint) error {
	category, err := s.catRepo.GetById(id)
	if err != nil {
		return fmt.Errorf("category not found")
	}

	if err := s.catRepo.DeleteCategory(category); err != nil {
		return fmt.Errorf("error to removing category")
	}

	return nil
}
