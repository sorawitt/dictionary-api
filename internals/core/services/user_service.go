package services

import (
	"dictionary-api/internals/core/ports"
	"errors"
)

type userService struct {
	userRepository ports.UserRepository
}

func NewUserService(repository ports.UserRepository) ports.UserService {
	return userService{userRepository: repository}
}

func (s userService) Login(email string, password string) error {
	err := s.userRepository.Login(email, password)
	if err != nil {
		return err
	}
	return nil
}
func (s userService) Register(email string, password string, passwordConfirmation string) error {
	if password != passwordConfirmation {
		return errors.New("the passwords are not equal")
	}
	err := s.userRepository.Register(email, password)
	if err != nil {
		return err
	}
	return nil
}
