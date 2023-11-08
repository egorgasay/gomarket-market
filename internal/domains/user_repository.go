package domains

import "go-rest-api/internal/model"

//go:generate go run github.com/vektra/mockery/v3 --name=IUserRepository
type IUserRepository interface {
	CreateUser(user model.User) error
}
