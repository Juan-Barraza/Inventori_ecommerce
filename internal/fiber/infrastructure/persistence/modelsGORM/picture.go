package modelsgorm

import "gorm.io/gorm"

type PicturesProduct struct {
	gorm.Model
	Link      string
	ProductId uint `gorm:"not null;index"`
}

type PicturesProductJson struct {
	ID        uint
	Link      string
	ProductId uint
}
