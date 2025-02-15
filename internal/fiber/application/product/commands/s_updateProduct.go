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

type UpdateProductService struct {
	productRep repositories.IProduct
	pictureRep repositories.IPicturesProductRepository
}

func NewUpdateProductService(productRep repositories.IProduct,
	pictureRep repositories.IPicturesProductRepository) *UpdateProductService {
	return &UpdateProductService{
		productRep: productRep,
		pictureRep: pictureRep,
	}
}

func (s *UpdateProductService) UpdateProduct(productData *domain.ProductJson, Prodid uint, files *multipart.Form) error {
	productExisting, err := s.productRep.GetById(Prodid)
	if err != nil {
		return fmt.Errorf("error to getting product")
	}

	if productData.Name != "" {
		productExisting.Name = productData.Name
	}
	if productData.Description != "" {
		productExisting.Description = productData.Description
	}
	if productData.Price != 0 {
		productExisting.Price = productData.Price
	}
	if productData.Stock != 0 {
		productExisting.Stock = productData.Stock
	}

	if err := s.productRep.UpdateProduct(productExisting); err != nil {
		return fmt.Errorf("error to updating product")
	}
	if files != nil {
		if err := s.updateFile(files, productExisting); err != nil {
			return err
		}
	}

	return nil
}

func (s *UpdateProductService) updateFile(files *multipart.Form, product *domain.Product) error {
	const maxSize int64 = 5 * 1024 * 1024
	picturesExisting, err := s.pictureRep.GetPicturesByProductId(product.ID)
	if err != nil {
		return fmt.Errorf("error to getting pictures")
	}

	for i := range picturesExisting {
		println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		pict := &picturesExisting[i]
		fmt.Printf("Eliminando archivo: %s\n", pict.Link)
		if err := os.Remove(pict.Link); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("error to remove file %s: %w", pict.Link, err)
		}
		println("eliminacion en media exitosa")

		if err := s.pictureRep.DeletePicture(pict); err != nil {
			return fmt.Errorf("error to remove image record: %w", err)
		}
		fmt.Println("file deleted")
		info, err := os.Stat(pict.Link)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Printf("El archivo %s no existe\n", pict.Link)
			} else {
				fmt.Printf("Error al acceder a %s: %v\n", pict.Link, err)
			}
		} else {
			fmt.Printf("El archivo %s existe: %v\n", pict.Link, info)
		}
		println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	}

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

		newPicture := &domain.PicturesProduct{
			ProductId: product.ID,
			Link:      filePath,
		}
		if err := s.pictureRep.CreatePicture(newPicture); err != nil {
			return fmt.Errorf("error al insertar la nueva imagen: %w", err)
		}

	}

	return nil
}
