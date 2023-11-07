package service

import (
	"errors"
	"go-rest-api/internal/domains/mocks"
	"go-rest-api/internal/model"
	"testing"
)

type sessionMock func(c *mocks.SessionUseCase)
type repositoryMock func(c *mocks.IUserRepository)

func Test_userService_SignUp(t *testing.T) {
	tests := []struct {
		name           string
		args           model.User
		sessionMock    sessionMock
		repositoryMock repositoryMock
		wantErr        error
	}{
		{
			name: "OK1",
			args: model.User{
				Username: "dima",
				Password: "test1",
				Session:  "ahsjufil12-fk",
			},
			sessionMock: func(c *mocks.SessionUseCase) {
				c.Mock.On("Generate").Return("ahsjufil12-fk", nil).Times(1)
			},
			repositoryMock: func(c *mocks.IUserRepository) {
				c.Mock.On("CreateUser", model.User{Username: "dima", Password: "test1", Session: "ahsjufil12-fk"}).Return(nil)
			},
			wantErr: nil,
		},
		{
			name: "OK1",
			args: model.User{
				Username: "dima",
				Password: "test1",
				Session:  "ahsjufil12-fk",
			},
			sessionMock: func(c *mocks.SessionUseCase) {
				c.Mock.On("Generate").Return("ahsjufil12-fk", nil).Times(1)
			},
			repositoryMock: func(c *mocks.IUserRepository) {
				c.Mock.On("CreateUser", model.User{Username: "dima", Password: "test1", Session: "ahsjufil12-fk"}).Return(errors.New("invalid")).Times(1)
			},
			wantErr: errors.New("invalid"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := mocks.NewIUserRepository(t)
			session := mocks.NewSessionUseCase(t)
			service := userService{
				database:       storage,
				sessionService: session,
			}
			_, err := service.SignUp(tt.args)
			if errors.Is(err, tt.wantErr) {
				t.Errorf("got %d, want %d", err, tt.wantErr)
			}
		})
	}
}
