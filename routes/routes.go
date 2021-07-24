package routes

import (
	"alterra_store/constants"
	"alterra_store/controllers"
	"alterra_store/middlewares"
	"alterra_store/validations"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	middlewares.LogMiddlewares(e)
	validations.CustomValidation(e)

	eAuth := e.Group("")
	eAuth.Use(middleware.JWT([]byte(constants.JWT_SECRET)))

	/** FEATURE USERS */
	e.POST("/api/v1/register", controllers.RegisterControllers)
	e.POST("/api/v1/login", controllers.LoginController)
	eAuth.GET("/api/v1/users", controllers.GetUserControllers)
	eAuth.GET("/api/v1/users/:userId", controllers.DetailUserControllers)
	eAuth.PUT("/api/v1/users", controllers.EditUserControllers)

	return e
}
