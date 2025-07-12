package handlers

import (
	"inventory/internal/fiber/application/client/commands"
	"inventory/internal/fiber/application/client/queries"
	domain "inventory/internal/fiber/domain/entities"
	"inventory/pkg/utils"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type ClienHandler struct {
	createclientService *commands.CreateClientCommandsService
	updateClientService *commands.UpdateClientCommandsService
	deleteCientService  *commands.DeleteClientCommandsService
	getAllClients       *queries.ClientQuerysService
}

func NewClientHandler(createclientService *commands.CreateClientCommandsService,
	updateClientService *commands.UpdateClientCommandsService,
	deleteCientService *commands.DeleteClientCommandsService,
	getAllClients *queries.ClientQuerysService,

) *ClienHandler {
	return &ClienHandler{
		createclientService: createclientService,
		updateClientService: updateClientService,
		deleteCientService:  deleteCientService,
		getAllClients:       getAllClients,
	}
}

func (h *ClienHandler) CreateClient(c fiber.Ctx) error {
	var clientData *domain.ClientJson
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

	if err := h.createclientService.CreateClient(client); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Client create"})
}

func (h *ClienHandler) GetAllClients(c fiber.Ctx) error {
	pagination := c.Locals("pagination").(*utils.Pagination)

	paginationResult, err := h.getAllClients.GetAll(pagination)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if paginationResult == nil {
		return c.Status(200).JSON(make([]domain.Client, 0))
	}

	return c.Status(200).JSON(paginationResult)
}

func (h *ClienHandler) ClientUpdate(c fiber.Ctx) error {
	clientId, err := strconv.Atoi(c.Params("id"))
	if (err != nil) || (clientId <= 0) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID invalido del cliente",
		})
	}
	var clientData domain.ClientJson
	if err := c.Bind().Body(&clientData); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "JSON not serializer",
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

	if err := h.updateClientService.UpdateClient(uint(clientId), client); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "update client",
	})
}

func (h *ClienHandler) ClientDelete(c fiber.Ctx) error {
	clientId, err := strconv.Atoi(c.Params("id"))
	if (err != nil) || (clientId <= 0) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID invalido del cliente",
		})
	}

	if err := h.deleteCientService.DeleteClient(uint(clientId)); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(204).JSON(fiber.Map{"message": "Client Delete sucessfully"})
}
