package modelsgorm

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
