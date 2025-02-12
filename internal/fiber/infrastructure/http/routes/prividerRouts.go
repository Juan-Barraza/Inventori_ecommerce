package routes

import (
	"inventory/internal/fiber/application/provider/commands"
	provider "inventory/internal/fiber/application/provider/queries"
	"inventory/internal/fiber/infrastructure/http/handlers"
	"inventory/internal/fiber/infrastructure/repository"
	"inventory/pkg"

	"github.com/gofiber/fiber/v3"
)

func SetProviderRoutes(apiV1 fiber.Router, db *pkg.Database) {
	paginationRep := repository.NewPaginationRepository(db)
	userRepo := repository.NewUserRepository(db)
	providerRepo := repository.NewProviderRepsoitor(db)
	createProvider := commands.NewCreateProviderService(providerRepo, userRepo)
	getProvider := provider.NewGetProviderService(providerRepo, paginationRep)
	updateProvider := commands.NewUpdateProviderService(providerRepo, userRepo)
	deleteProvider := commands.NewDeleteProviderService(providerRepo, userRepo)
	handlerProvider := handlers.NewHandlerProvider(createProvider,
		getProvider, updateProvider, deleteProvider)

	apiV1.Post("/providers", handlerProvider.CreateProvider)
	apiV1.Get("/providers", handlerProvider.GetAllProvider)
	apiV1.Put("/providers/:id", handlerProvider.ProviderUpdate)
	apiV1.Delete("/providers/:id", handlerProvider.DeleteProvider)
}
