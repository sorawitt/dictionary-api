package server

import (
	"dictionary-api/internals/core/ports"
	"log"

	"github.com/gofiber/fiber/v2"
)

type server struct {
	userHandlers ports.UserHandlers
}

func NewServer(userHandlers ports.UserHandlers) server {
	return server{
		userHandlers: userHandlers,
	}
}

func (s server) Initialize() {
	app := fiber.New()
	v1 := app.Group("/v1")

	userRoutes := v1.Group("/user")
	userRoutes.Post("/login", s.userHandlers.Login)
	userRoutes.Post("/register", s.userHandlers.Register)

	err := app.Listen(":5050")
	if err != nil {
		log.Fatal(err)
	}
}
