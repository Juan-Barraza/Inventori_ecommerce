package routes

import (
	"inventory/internal/fiber/application"
	"inventory/internal/fiber/infrastructure/http/handlers"
	"inventory/internal/fiber/infrastructure/persistence/repository"
	"inventory/pkg"

	"github.com/gofiber/fiber/v3"
)

func SetClientRoutes(apiV1 fiber.Router, db *pkg.Database) {
	userRepo := repository.NewUserRepository(db)
	clientRepo := repository.NewClientRepository(db)
	clientService := application.NewClientService(clientRepo, userRepo)
	clientHandler := handlers.NewClientHandler(
		clientService,
	)

	apiV1.Post("/clients", clientHandler.CreateClient)
	apiV1.Get("/clients", clientHandler.GetAllClients)
}
