package ports

import "github.com/gofiber/fiber/v2"

type UserService interface {
	Login(email string, password string) error
	Register(email string, password string, passwordConfirmation string) error
}

type UserRepository interface {
	Login(email string, password string) error
	Register(email string, password string) error
}

type UserHandlers interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}
