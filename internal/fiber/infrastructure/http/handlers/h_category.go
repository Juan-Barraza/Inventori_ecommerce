package handlers

import (
	"inventory/internal/fiber/application/category/commands"
	"inventory/internal/fiber/application/category/queries"
	domain "inventory/internal/fiber/domain/entities"
	"inventory/pkg/utils"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v3"
)

type CategoryHandler struct {
	createCategoryService   *commands.CreateCategoryService
	getAllCategoriesService *queries.GetAllCategoriesService
	deleteCategoryService   *commands.DeleteCategoryService
}

func NewCategoryHandler(
	createCategoryService *commands.CreateCategoryService,
	getAllCategoriesService *queries.GetAllCategoriesService,
	deleteCategoryService *commands.DeleteCategoryService) *CategoryHandler {
	return &CategoryHandler{
		createCategoryService:   createCategoryService,
		getAllCategoriesService: getAllCategoriesService,
		deleteCategoryService:   deleteCategoryService}
}

func (h *CategoryHandler) CreateCategory(c fiber.Ctx) error {
	var categoryData *domain.CategoryJson

	if err := c.Bind().Body(&categoryData); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "error to parse JSON",
		})
	}

	category := &domain.Category{
		CategoryName: categoryData.CategoryName,
		Description:  categoryData.Description,
	}

	if err := h.createCategoryService.CreateCategory(category); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Category created sucessfully"})
}

func (h *CategoryHandler) GetAllCategories(c fiber.Ctx) error {
	pagination := c.Locals("pagination").(*utils.Pagination)
	categoriesPagination, err := h.getAllCategoriesService.GetAllCategories(pagination)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if categoriesPagination == nil {
		return c.Status(200).JSON(make([]domain.Category, 0))
	}

	return c.Status(200).JSON(categoriesPagination)
}

func (h *CategoryHandler) DeleteCategory(c fiber.Ctx) error {
	categoryID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid ID parameter",
		})
	}
	if err := h.deleteCategoryService.DeleteCategory(uint(categoryID)); err != nil {
		if strings.Contains(err.Error(), "category not found") {
			return c.Status(404).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(204)

}
