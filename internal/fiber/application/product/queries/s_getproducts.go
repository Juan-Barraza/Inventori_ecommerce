package queries

import (
	"fmt"
	domain "inventory/internal/fiber/domain/entities"
	"inventory/internal/fiber/domain/repositories"
	"inventory/internal/fiber/infrastructure/repository"
	"inventory/pkg/utils"
)

type GetProductsService struct {
	productRep    repositories.IProduct
	paginationRep *repository.PaginationRepository
}

func NewGetProductsService(productRep repositories.IProduct,
	paginationRep *repository.PaginationRepository) *GetProductsService {
	return &GetProductsService{
		productRep:    productRep,
		paginationRep: paginationRep}
}

func (s *GetProductsService) GetProduct(category string, providerID uint, pagination *utils.Pagination) (*utils.Pagination, error) {
	query, products, err := s.productRep.GetAllProducts(category, providerID)
	if err != nil {
		return nil, fmt.Errorf("error to get products")
	}

	paginationResult, err := s.paginationRep.GetPaginatedResults(query, pagination, &products)
	if err != nil {
		return nil, fmt.Errorf("error to make pagination")
	}
	if products == nil {
		products = []domain.Product{}
	}
	var productsJson []domain.ProductJson

	for _, product := range products {
		productsJson = append(productsJson, *domain.ToProduct(&product))
	}

	paginationResult.Data = productsJson
	return paginationResult, nil

}
