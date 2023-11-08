package service

import (
	"go-rest-api/internal/domains"
	"go-rest-api/internal/model"
	"go-rest-api/internal/validator"
)

type userService struct {
	database       domains.IUserRepository
	validator      validator.IUserValidator
	sessionService domains.SessionUseCase
}

func NewUserUseCase(database domains.IUserRepository, validator validator.IUserValidator, sessionService domains.SessionUseCase) domains.UserUseCase {
	return &userService{database, validator, sessionService}
}

func (s *userService) SignUp(user model.User) (string, error) {
	//if err := s.validator.UserValidate(user); err != nil {
	//	return "", err
	//}
	//hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	//if err != nil {
	//	return err
	//}
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
