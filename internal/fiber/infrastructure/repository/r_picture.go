package repository

import (
	domain "inventory/internal/fiber/domain/entities"
	"inventory/pkg"
)

type PictureRepository struct {
	db *pkg.Database
}

func NewPictureRepository(db *pkg.Database) *PictureRepository {
	return &PictureRepository{db: db}
}

func (r *PictureRepository) CreatePicture(picture *domain.PicturesProduct) error {
	return r.db.DB.Create(picture).Error
}

func (r *PictureRepository) GetPictures() ([]domain.PicturesProduct, error) {
	var pictures []domain.PicturesProduct

	if err := r.db.DB.Model(&domain.PicturesProduct{}).
		Find(&pictures).Error; err != nil {
		return nil, err
	}

	return pictures, nil
}

func (r *PictureRepository) GetByID(id uint) (*domain.PicturesProduct, error) {
	var picture *domain.PicturesProduct

	if err := r.db.DB.Model(&domain.PicturesProduct{}).
		First(&picture, id).Error; err != nil {
		return nil, err
	}

	return picture, nil
}

func (r *PictureRepository) UpdatePicture(picture *domain.PicturesProduct) error {
	return r.db.DB.Model(&picture).Where("id = ?", picture.ID).Updates(picture).Error
}

func (r *PictureRepository) DeletePicture(picture *domain.PicturesProduct) error {
	return r.db.DB.Unscoped().Delete(picture).Error
}

func (r *PictureRepository) GetPicturesByProductId(productId uint) ([]domain.PicturesProduct, error) {
	var pictures []domain.PicturesProduct
	if err := r.db.DB.Model(&domain.PicturesProduct{}).
		Where("product_id = ?", productId).
		Find(&pictures).Error; err != nil {
		return nil, err
	}

	return pictures, nil
}
