package routes

import (
	"inventory/internal/fiber/application/provider/commands"
	"inventory/internal/fiber/infrastructure/http/handlers"
	"inventory/internal/fiber/infrastructure/persistence/repository"
	"inventory/pkg"

	"github.com/gofiber/fiber/v3"
)

func SetProviderRoutes(apiV1 fiber.Router, db *pkg.Database) {
	userRepo := repository.NewUserRepository(db)
	providerRepo := repository.NewProviderRepsoitor(db)
	createProvider := commands.NewCreateProviderService(providerRepo, userRepo)
	handlerProvider := handlers.NewHandlerProvider(createProvider)

	apiV1.Post("/providers", handlerProvider.CreateProvider)
}
