package routes

import (
	"inventory/internal/fiber/application/product/commands"
	"inventory/internal/fiber/application/product/queries"
	"inventory/internal/fiber/infrastructure/http/handlers"
	"inventory/internal/fiber/infrastructure/repository"
	"inventory/pkg"

	"github.com/gofiber/fiber/v3"
)

func SetProductsRoutes(apiV1 fiber.Router, db *pkg.Database) {
	paginationRepo := repository.NewPaginationRepository(db)
	productRepository := repository.NewProductRespository(db)
	pictureRepository := repository.NewPictureRepository(db)
	createProdService := commands.NewAddProductService(productRepository, pictureRepository)
	getProductService := queries.NewGetProductsService(productRepository, paginationRepo)
	deleteProductService := commands.NewDeleteProductService(productRepository)
	updateProductService := commands.NewUpdateProductService(productRepository, pictureRepository)
	productHandler := handlers.NewProductHandler(createProdService,
		getProductService,
		deleteProductService,
		updateProductService)

	apiV1.Post("/products", productHandler.CreateProduct)
	apiV1.Get("/products", productHandler.GetProduct)
	apiV1.Delete("/products/:id", productHandler.DeleteProduct)
	apiV1.Put("/products/:id", productHandler.UpdateProduct)
}
