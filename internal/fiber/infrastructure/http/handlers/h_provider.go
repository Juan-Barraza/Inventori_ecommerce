package handlers

import (
	"inventory/internal/fiber/application/provider/commands"
	provider "inventory/internal/fiber/application/provider/queries"
	domain "inventory/internal/fiber/domain/entities"
	"inventory/pkg/utils"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type ProviderHandler struct {
	createProviderservice *commands.CreateProviderService
	getallProviderService *provider.GetProviderService
	updateProviderService *commands.UpdateProviderService
	deleteProviderService *commands.DeleteProviderService
}

func NewHandlerProvider(createProvider *commands.CreateProviderService,
	getallProviderService *provider.GetProviderService,
	updateProviderService *commands.UpdateProviderService,
	deleteProviderService *commands.DeleteProviderService) *ProviderHandler {
	return &ProviderHandler{
		createProviderservice: createProvider,
		getallProviderService: getallProviderService,
		updateProviderService: updateProviderService,
		deleteProviderService: deleteProviderService,
	}
}

func (h *ProviderHandler) CreateProvider(c fiber.Ctx) error {
	var providerData domain.ProviderJson
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

func (h *ProviderHandler) GetAllProvider(c fiber.Ctx) error {
	pagination := c.Locals("pagination").(*utils.Pagination)

	providers, err := h.getallProviderService.GetALL(pagination)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if providers == nil {
		return c.Status(200).JSON(make([]domain.Provider, 0))
	}

	return c.Status(200).JSON(providers)
}

func (h *ProviderHandler) ProviderUpdate(c fiber.Ctx) error {
	providerId, err := strconv.Atoi(c.Params("id"))
	if (err != nil) || (providerId <= 0) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID not valid",
		})
	}
	var providerJson domain.ProviderJson
	if err = c.Bind().Body(&providerJson); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "error to parse JSON",
		})
	}

	provider := &domain.Provider{
		Name:          providerJson.Name,
		Address:       providerJson.Address,
		TypeOfProduct: providerJson.TypeOfProduct,
		PhoneNumber:   providerJson.PhoneNumber,
		User: domain.User{
			Email:    providerJson.Email,
			Password: providerJson.Password,
		},
	}

	err = h.updateProviderService.UpdateProvider(uint(providerId), provider)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Provider update sucessfully"})

}

func (h *ProviderHandler) DeleteProvider(c fiber.Ctx) error {
	providerId, err := strconv.Atoi(c.Params("id"))
	if err != nil || providerId == 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "id invalid",
		})
	}

	if err = h.deleteProviderService.DeleteProvider(uint(providerId)); err != nil {
		if err.Error() == "provider not found" || err.Error() == "user not found" {
			return c.Status(404).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	return c.Status(204).JSON(fiber.Map{"message": "Delete provvider sucessfully"})
}
