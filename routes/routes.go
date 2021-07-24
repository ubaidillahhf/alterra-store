package routes

import (
	"alterra_store/constants"
	"alterra_store/controllers"
	"alterra_store/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	middlewares.LogMiddlewares(e)
	eAuth := e.Group("")
	eAuth.Use(middleware.JWT([]byte(constants.JWT_SECRET)))

	/** FEATURE USERS */
	e.POST("/api/v1/register", controllers.RegisterControllers)
	e.POST("/api/v1/login", controllers.LoginController)
	eAuth.GET("/api/v1/users", controllers.GetUserControllers)
	eAuth.GET("/api/v1/users/:userId", controllers.DetailUserControllers)

	return e
}
