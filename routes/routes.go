package routes

import (
	"alterra_store/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.POST("/api/v1/register", controllers.CreateUserControllers)
	e.POST("/api/v1/login", controllers.LoginController)
	return e
}
