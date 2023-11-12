package service

import (
	"fmt"
	"go-rest-api/internal/domains"
	"go-rest-api/internal/model"
	"go-rest-api/internal/validator"
)

type Service struct {
	database       domains.IRepository
	validator      validator.IUserValidator
	sessionService domains.SessionService
}

func NewUserUseCase(database domains.IRepository, validator validator.IUserValidator, sessionService domains.SessionService) domains.Service {
	return &Service{database, validator, sessionService}
}

func (s *Service) SignUp(user model.User) (string, error) {
	uuid, err := s.sessionService.Generate()
	if err != nil {
		return "", err
	}
	newUser := model.User{Username: user.Username, Password: user.Password, Session: uuid}
	if err := s.database.CreateUser(newUser); err != nil {
		return "", err
	}
	return uuid, nil
}

func (s *Service) Login(user model.User) error {
	data := model.User{}
	if err := s.database.GetUserByUsername(&data, user.Username); err != nil {
		return fmt.Errorf("invalid data for login")
	}
	return nil
}
