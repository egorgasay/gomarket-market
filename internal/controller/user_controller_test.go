package controller

import (
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"go-rest-api/internal/domains/mocks"
	"go-rest-api/internal/model"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type serviceMock func(c *mocks.UserUseCase)
type sessionMock func(c *mocks.SessionUseCase)

func Test_UserController_SignUp(t *testing.T) {
	type jso struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	tests := []struct {
		name        string
		body        jso
		sessionMock sessionMock
		serviceMock serviceMock
		wantCode    int
	}{
		{
			name: "OK1",
			body: jso{
				Login:    "dima",
				Password: "test1",
			},
			sessionMock: func(c *mocks.SessionUseCase) {
				c.Mock.On("Generate").Return("ahsjufil12-fk", nil)
			},
			serviceMock: func(c *mocks.UserUseCase) {
				c.Mock.On("SignUp", model.User{Username: "dima", Password: "test1", Session: "ahsjufil12-fk"}).Return("ahsjufil12-fk", nil).Times(1)
			},
			wantCode: http.StatusOK,
		},
		{
			name: "BAD",
			body: jso{
				Login: "dima",
			},
			sessionMock: func(c *mocks.SessionUseCase) {
				c.Mock.On("Generate").Return("ahsjufil12-fk", nil)
			},
			serviceMock: func(c *mocks.UserUseCase) {
				user := model.User{Username: "dima", Password: "", Session: "ahsjufil12-fk"}
				c.Mock.On("SignUp", user).Return("", errors.New("invalid")).Times(1)
			},
			wantCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			service := mocks.NewUserUseCase(t)
			sessionMock := mocks.NewSessionUseCase(t)
			h := NewUserController(service)
			tt.serviceMock(service)
			tt.sessionMock(sessionMock)
			b, err := json.Marshal(tt.body)
			if err != nil {
				return
			}
			path := "/v1/user/register"
			e.POST(path, h.SignUp)

			w := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodPost, path, strings.NewReader(string(b)))
			// создаём новый Recorder

			e.ServeHTTP(w, request)

			if w.Code != tt.wantCode {
				t.Errorf("got %d, want %d", w.Code, tt.wantCode)
			}
		})
	}
}
