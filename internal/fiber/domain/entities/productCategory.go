package domain

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryName string
	Description  string
}

type CategoryJson struct {
	ID           uint
	CategoryName string
	Description  string
}

func ToCategory(c *Category) *CategoryJson {
	return &CategoryJson{
		ID:           c.ID,
		CategoryName: c.CategoryName,
		Description:  c.Description,
	}
}
