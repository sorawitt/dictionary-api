package handlers

import (
	"dictionary-api/internals/core/ports"

	"github.com/gofiber/fiber/v2"
)

type user struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type userRegister struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

type userHandlers struct {
	userService ports.UserService
}

func NewUserHandlers(userService ports.UserService) ports.UserHandlers {
	return userHandlers{userService: userService}
}

func (h userHandlers) Login(c *fiber.Ctx) error {
	u := new(user)
	if err := c.BodyParser(u); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	err := h.userService.Login(u.Email, u.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return nil
}

func (h userHandlers) Register(c *fiber.Ctx) error {
	u := new(userRegister)
	if err := c.BodyParser(u); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	err := h.userService.Register(u.Email, u.Password, u.PasswordConfirmation)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return nil
}
