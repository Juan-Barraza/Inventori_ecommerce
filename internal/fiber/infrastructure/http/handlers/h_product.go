package handlers

import (
	"inventory/internal/fiber/application/product/commands"
	"inventory/internal/fiber/application/product/queries"
	domain "inventory/internal/fiber/domain/entities"
	"inventory/pkg/utils"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v3"
)

type ProductHandler struct {
	createProduct *commands.AddProductService
	getProducts   *queries.GetProductsService
	deleteProduct *commands.DeleteProductService
	updateProduct *commands.UpdateProductService
}

func NewProductHandler(createProduct *commands.AddProductService,
	getProducts *queries.GetProductsService,
	deleteProduct *commands.DeleteProductService,
	updateProduct *commands.UpdateProductService) *ProductHandler {
	return &ProductHandler{
		createProduct: createProduct,
		getProducts:   getProducts,
		deleteProduct: deleteProduct,
		updateProduct: updateProduct,
	}
}

func (h *ProductHandler) CreateProduct(c fiber.Ctx) error {
	name := c.FormValue("name")
	description := c.FormValue("description")
	priceStr := c.FormValue("price")
	stockStr := c.FormValue("stock")
	categoryIdStr := c.FormValue("categoryId")
	providerIdStr := c.FormValue("providerId")

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Price must be a valid number"})
	}
	stock, err := strconv.Atoi(stockStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Stock must be a valid integer"})
	}
	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "CategoryId must be a valid integer"})
	}
	providerId, err := strconv.Atoi(providerIdStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ProviderId must be a valid integer"})
	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Error processing multipart form data"})
	}


	product := &domain.Product{
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
		CategoryId:  uint(categoryId),
		ProviderId:  uint(providerId),
	}
	if err := h.createProduct.CreateProduct(product, form); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "creation sucessfully"})
}

func (h *ProductHandler) GetProduct(c fiber.Ctx) error {
	pagination := c.Locals("pagination").(*utils.Pagination)
	category := c.Query("category")
	providerIdStr := c.Query("providerId")
	var provider int
	var err error

	if providerIdStr != "" {
		provider, err = strconv.Atoi(providerIdStr)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "error to parse id",
			})
		}

	}

	productsPagination, err := h.getProducts.GetProduct(category, uint(provider), pagination)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(productsPagination)

}

func (h *ProductHandler) DeleteProduct(c fiber.Ctx) error {
	idProduct, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id parameter",
		})
	}

	if err := h.deleteProduct.DeleteProduct(uint(idProduct)); err != nil {
		if strings.Contains(err.Error(), "error to getting product to remove") {
			return c.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(204)
}

func (h *ProductHandler) UpdateProduct(c fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id paremeter",
		})
	}
	name := c.FormValue("name")
	description := c.FormValue("description")
	priceStr := c.FormValue("price")
	stockStr := c.FormValue("stock")
	var price float64
	var stock int

	if priceStr != "" {
		price, err = strconv.ParseFloat(priceStr, 64)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Price must be a valid number"})
		}
	}
	if stockStr != "" {
		stock, err = strconv.Atoi(stockStr)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Stock must be a valid integer"})
		}
	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Error processing multipart form data"})
	}

	productData := &domain.ProductJson{
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
	}

	if err := h.updateProduct.UpdateProduct(productData, uint(productId), form); err != nil {
		if strings.Contains(err.Error(), "error to getting product") {
			return c.Status(400).JSON(fiber.Map{
				"error": "Product not found",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{"message": "updated sucessfully"})
}
