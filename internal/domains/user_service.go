package domains

import "go-rest-api/internal/model"

//go:generate go run github.com/vektra/mockery/v3 --name=UserUseCase
type UserUseCase interface {
	SignUp(user model.User) (string, error)
}
