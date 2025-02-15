package repositories

import domain "inventory/internal/fiber/domain/entities"

type IPicturesProductRepository interface {
	CreatePicture(picture *domain.PicturesProduct) error
	GetPictures() ([]domain.PicturesProduct, error)
	GetByID(id uint) (*domain.PicturesProduct, error)
	GetPicturesByProductId(id uint) ([]domain.PicturesProduct, error)
	UpdatePicture(picture *domain.PicturesProduct) error
	DeletePicture(picture *domain.PicturesProduct) error
}
