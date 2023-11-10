package controller

import (
	"github.com/labstack/echo/v4"
	"go-rest-api/internal/domains"
	"go-rest-api/internal/errors"
	"go-rest-api/internal/model"
	"net/http"
	"time"
)

type IUserController interface {
	SignUp(c echo.Context) error
}

type userController struct {
	userService domains.Service
}

func NewUserController(userService domains.Service) IUserController {
	return &userController{userService}
}

func (s *userController) SignUp(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return errors.Handler(c, err)
	}

	uuid, err := s.userService.SignUp(user)
	if err != nil {
		return errors.Handler(c, err)
	}
	cookie := http.Cookie{}
	cookie.Name = "token"
	cookie.Value = uuid
	cookie.Expires = time.Now().Add(56 * time.Hour)
	c.SetCookie(&cookie)
	c.Response().Status = http.StatusOK
	return nil
}
