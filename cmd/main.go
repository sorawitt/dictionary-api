package main

import (
	"dictionary-api/internals/core/services"
	"dictionary-api/internals/handlers"
	"dictionary-api/internals/repositories"
	"dictionary-api/internals/server"
	"log"
)

func main() {
	mongoConn := ""

	userRepository, err := repositories.NewUserRepository(mongoConn)
	if err != nil {
		log.Fatal(err)
	}
	userService := services.NewUserService(userRepository)
	userHandlers := handlers.NewUserHandlers(userService)

	httpServer := server.NewServer(userHandlers)
	httpServer.Initialize()

}
