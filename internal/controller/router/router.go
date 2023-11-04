package router

import (
	"github.com/labstack/echo/v4"
	controller2 "go-rest-api/internal/controller"
)

func NewRouter(uc controller2.IUserController) *echo.Echo {
	e := echo.New()
	//e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//	AllowOrigins: []string{"http://localhost:8080", os.Getenv("FE_URL")},
	//	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
	//		echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
	//	AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
	//	AllowCredentials: true,
	//}))
	//e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
	//	CookiePath:     "/",
	//	CookieDomain:   os.Getenv("API_DOMAIN"),
	//	CookieHTTPOnly: true,
	//	CookieSameSite: http.SameSiteNoneMode,
	//}))
	e.POST("/v1/user/register", uc.SignUp)
	return e
}
