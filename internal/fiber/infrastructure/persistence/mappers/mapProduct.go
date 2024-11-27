package mappers

import (
	domain "inventory/internal/fiber/domain/models"
	modelsgorm "inventory/internal/fiber/infrastructure/persistence/modelsGORM"
)

func ToProductGorm(p *domain.Product) *modelsgorm.Product {
	images := make([]modelsgorm.PicturesProduct, len(p.Images))
	for i, img := range p.Images {
		images[i] = modelsgorm.PicturesProduct{
			Link:      img.Link,
			ProductId: img.ProductId,
		}
	}
	var product = &modelsgorm.Product{
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Images:      images,
		Stock:       p.Stock,
		CategoryId:  p.CategoryId,
		ProviderId:  p.CategoryId,
	}
	return product
}

func FromProductGorm(p *modelsgorm.Product) *domain.Product {
	images := make([]domain.PicturesProduct, len(p.Images))
	for i, img := range p.Images {
		images[i] = domain.PicturesProduct{
			Link:      img.Link,
			ProductId: img.ProductId,
		}
	}
	var product = &domain.Product{
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Images:      images,
		Stock:       p.Stock,
		CategoryId:  p.CategoryId,
		ProviderId:  p.CategoryId,
	}

	return product
}

func ToCategoryGorm(c *domain.Category) modelsgorm.Category {
	return modelsgorm.Category{
		CategoryName: c.CategoryName,
		Description:  c.Description,
	}
}
