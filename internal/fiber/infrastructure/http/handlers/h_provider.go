package handlers

import (
	"inventory/internal/fiber/application/provider/commands"
	domain "inventory/internal/fiber/domain/entities"
	modelsgorm "inventory/internal/fiber/infrastructure/persistence/modelsGORM"

	"github.com/gofiber/fiber/v3"
)

type ProviderHandler struct {
	createProviderservice *commands.CreateProviderService
}

func NewHandlerProvider(createProvider *commands.CreateProviderService) *ProviderHandler {
	return &ProviderHandler{
		createProviderservice: createProvider,
	}
}

func (h *ProviderHandler) CreateProvider(c fiber.Ctx) error {
	var providerData modelsgorm.ProviderJson
	if err := c.Bind().Body(&providerData); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "error to parse JSON",
		})
	}

	provider := &domain.Provider{
		Name:          providerData.Name,
		Address:       providerData.Address,
		PhoneNumber:   providerData.PhoneNumber,
		TypeOfProduct: providerData.TypeOfProduct,
		User: domain.User{
			Email:    providerData.Email,
			Password: providerData.Password,
		},
	}

	if err := h.createProviderservice.CreateProvider(provider); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Provider create sucessfully"})
}
