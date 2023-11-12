package router

import (
	"github.com/labstack/echo/v4"
	controller2 "go-rest-api/internal/controller"
)

func NewRouter(uc controller2.IUserController) *echo.Echo {
	e := echo.New()
	e.POST("/v1/user/register", uc.SignUp)
	e.POST("/v1/user/login", uc.Login)
	return e
}
