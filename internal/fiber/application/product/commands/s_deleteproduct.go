package commands

import (
	"fmt"
	"inventory/internal/fiber/domain/repositories"
)

type DeleteProductService struct {
	productRepo repositories.IProduct
}

func NewDeleteProductService(productRepo repositories.IProduct) *DeleteProductService {
	return &DeleteProductService{productRepo: productRepo}
}

func (s *DeleteProductService) DeleteProduct(id uint) error {
	product, err := s.productRepo.GetById(id)
	if err != nil {
		return fmt.Errorf("error to getting product to remove")
	}
	println("Product ID:", product.ID)

	if err := s.productRepo.DeleteProduct(product); err != nil {
		return fmt.Errorf("error to removing product")
	}
	println("delete sucessfully")

	return nil
}
