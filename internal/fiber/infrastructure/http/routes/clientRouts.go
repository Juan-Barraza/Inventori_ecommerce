package routes

import (
	"inventory/internal/fiber/application/client/commands"
	"inventory/internal/fiber/application/client/queries"
	"inventory/internal/fiber/infrastructure/http/handlers"
	"inventory/internal/fiber/infrastructure/repository"
	"inventory/pkg"

	"github.com/gofiber/fiber/v3"
)

func SetClientRoutes(apiV1 fiber.Router, db *pkg.Database) {
	userRepo := repository.NewUserRepository(db)
	clientRepo := repository.NewClientRepository(db)
	paginationRepo := repository.NewPaginationRepository(db)
	createClientS := commands.NewClientCommandsService(clientRepo, userRepo)
	updateClientS := commands.NewUpdateClientCommandsService(clientRepo, userRepo)
	deleteClientS := commands.NewDeleteClientCommandsService(clientRepo, userRepo)
	getAllClientS := queries.NewClientQuerysService(clientRepo, paginationRepo)
	clientHandler := handlers.NewClientHandler(
		createClientS,
		updateClientS,
		deleteClientS,
		getAllClientS,
	)

	apiV1.Post("/clients", clientHandler.CreateClient)
	apiV1.Get("/clients", clientHandler.GetAllClients)
	apiV1.Put("/clients/:id", clientHandler.ClientUpdate)
	apiV1.Delete("/clients/:id", clientHandler.ClientDelete)
}
