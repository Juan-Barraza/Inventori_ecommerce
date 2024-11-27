package handlers

import (
	"inventory/internal/fiber/application"
	domain "inventory/internal/fiber/domain/models"
	modelsgorm "inventory/internal/fiber/infrastructure/persistence/modelsGORM"

	"github.com/gofiber/fiber/v3"
)

type ClienHandler struct {
	clientService *application.ClientService
}

func NewClientHandler(clientServ *application.ClientService) *ClienHandler {
	return &ClienHandler{
		clientService: clientServ,
	}
}

func (h *ClienHandler) CreateClient(c fiber.Ctx) error {
	var clientData *modelsgorm.ClientJson
	if err := c.Bind().Body(&clientData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "JSON not serializable",
		})
	}

	client := &domain.Client{
		Name:           clientData.Name,
		LastName:       clientData.LastName,
		TypeDocument:   clientData.TypeDocument,
		DocumentNumber: clientData.DocumentNumber,
		PhoneNumber:    clientData.PhoneNumber,
		Address:        clientData.Address,
		User: domain.User{
			Email:    clientData.Email,
			Password: clientData.Password,
		},
	}

	if err := h.clientService.CreateClient(client); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Client create"})
}

func (h *ClienHandler) GetAllClients(c fiber.Ctx) error {
	clients, err := h.clientService.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "clients no obtains",
		})
	}

	return c.Status(200).JSON(clients)
}
