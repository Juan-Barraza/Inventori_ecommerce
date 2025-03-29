package routes

import (
	"inventory/internal/fiber/application/category/commands"
	"inventory/internal/fiber/application/category/queries"
	"inventory/internal/fiber/infrastructure/http/handlers"
	"inventory/internal/fiber/infrastructure/repository"
	"inventory/pkg"

	"github.com/gofiber/fiber/v3"
)

func SetCategoryRoutes(apiV1 fiber.Router, db *pkg.Database) {
	categoryRepo := repository.NewCategoryRepository(db)
	paginationRepo := repository.NewPaginationRepository(db)
	createCategoryS := commands.NewCreateCategoryService(categoryRepo)
	getAllCategoryS := queries.NewGetAllCategoriesService(categoryRepo, paginationRepo)
	deleteCategoryS := commands.NewDeleteCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(createCategoryS, getAllCategoryS, deleteCategoryS)

	apiV1.Post("/categories", categoryHandler.CreateCategory)
	apiV1.Get("/categories", categoryHandler.GetAllCategories)
	apiV1.Delete("/categories/:id", categoryHandler.DeleteCategory)

}
