package service

import (
	"go-rest-api/internal/model"
	"go-rest-api/internal/repository"
	"go-rest-api/internal/validator"
)

type IUserService interface {
	SignUp(user model.User) (string, error)
}

type userService struct {
	database       repository.IUserRepository
	validator      validator.IUserValidator
	sessionService SessionService
}

func NewUserUsecase(database repository.IUserRepository, validator validator.IUserValidator, sessionService SessionService) IUserService {
	return &userService{database, validator, sessionService}
}

func (s *userService) SignUp(user model.User) (string, error) {
	if err := s.validator.UserValidate(user); err != nil {
		return "", err
	}
	//hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	//if err != nil {
	//	return err
	//}
	uuid, err := s.sessionService.Generate()
	if err != nil {
		return "", err
	}
	newUser := model.User{Username: user.Username, Password: user.Password, Session: uuid}
	if err := s.database.CreateUser(&newUser); err != nil {
		return "", err
	}
	return uuid, nil
}
