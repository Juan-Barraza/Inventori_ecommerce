package routes

import (
	"inventory/internal/fiber/application"
	"inventory/internal/fiber/infrastructure/http/handlers"
	"inventory/internal/fiber/infrastructure/repository"
	"inventory/pkg"

	"github.com/gofiber/fiber/v3"
)

func SetUserRoutes(apiV1 fiber.Router, db *pkg.Database) {
	userRepo := repository.NewUserRepository(db)
	paginationRep := repository.NewPaginationRepository(db)
	userService := application.NewUserService(userRepo, paginationRep)
	userHandler := handlers.NewUserHandler(
		userService,
	)

	apiV1.Post("/users", userHandler.Register)
	apiV1.Get("/users", userHandler.GetAllUsers)
}
