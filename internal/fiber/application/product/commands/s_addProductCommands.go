package commands

import (
	"fmt"
	domain "inventory/internal/fiber/domain/entities"
	"inventory/internal/fiber/domain/repositories"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type AddProductService struct {
	productRep repositories.IProduct
	pictureRep repositories.IPicturesProductRepository
}

func NewAddProductService(productRep repositories.IProduct,
	pictureRep repositories.IPicturesProductRepository) *AddProductService {
	return &AddProductService{productRep: productRep,
		pictureRep: pictureRep}
}

func (s *AddProductService) CreateProduct(prod *domain.Product, files *multipart.Form) error {
	const maxSize int64 = 5 * 1024 * 1024
	if err := s.productRep.ValidateUniqueProduct(prod); err != nil {
		return err
	}

	if err := s.productRep.AddProduct(prod); err != nil {
		return fmt.Errorf("error to create product")
	}

	if files != nil {
		imageFiles := files.File["images"]
		for _, fileHeader := range imageFiles {
			if fileHeader.Size > maxSize {
				return fmt.Errorf("the file %s exceeds the size of 5 MB", fileHeader.Filename)
			}

			ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
			if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
				return fmt.Errorf("the file %s has an unallowed extension (%s)", fileHeader.Filename, ext)
			}

			file, err := fileHeader.Open()
			if err != nil {
				return fmt.Errorf("error opening file: %w", err)
			}

			filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), fileHeader.Filename)
			filePath := filepath.Join("/app/media", filename)

			out, err := os.Create(filePath)
			if err != nil {
				file.Close()
				return fmt.Errorf("error creating file on disk: %w", err)
			}

			if _, err := io.Copy(out, file); err != nil {
				out.Close()
				file.Close()
				return fmt.Errorf("error saving file: %w", err)
			}

			if err := out.Close(); err != nil {
				file.Close()
				return fmt.Errorf("error closing output file: %w", err)
			}
			if err := file.Close(); err != nil {
				return fmt.Errorf("error closing input file: %w", err)
			}

			picture := &domain.PicturesProduct{
				Link:      filePath,
				ProductId: prod.ID,
			}

			if err := s.pictureRep.CreatePicture(picture); err != nil {
				return fmt.Errorf("error to create picture in database")
			}
		}
	}

	return nil
}
