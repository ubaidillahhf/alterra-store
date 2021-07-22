package routes

import (
	"alterra_store/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.POST("/users", controllers.CreateUsersControllers)
	return e
}
