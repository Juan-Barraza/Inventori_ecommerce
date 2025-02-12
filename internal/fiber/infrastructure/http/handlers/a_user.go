package handlers

import (
	"inventory/internal/fiber/application"
	domain "inventory/internal/fiber/domain/entities"
	"inventory/pkg/utils"

	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	userService *application.UserService
}

func NewUserHandler(service *application.UserService) *UserHandler {
	return &UserHandler{userService: service}
}

func (h *UserHandler) Register(c fiber.Ctx) error {
	var userData domain.UserGormJson

	if err := c.Bind().Body(&userData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error al crear user",
		})
	}
	password, err := utils.HashPassword(userData.Password)
	if err != nil {
		return err
	}
	user := domain.User{
		Email:    userData.Email,
		Password: password,
	}

	if err := h.userService.RegisterUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "no create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"menssage": "create user"})
}

func (h *UserHandler) GetAllUsers(c fiber.Ctx) error {
	users, err := h.userService.GelAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Users not conted",
		})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}
