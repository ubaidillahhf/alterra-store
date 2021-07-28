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
	e.Pre(middleware.RemoveTrailingSlash())
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
	eAuth.DELETE("/api/v1/users", controllers.DeleteUserControllers)

	/** FEATURE PRODUCT CATEGORY */
	eAuth.POST("/api/v1/product-category", controllers.CreateCategoryControllers)
	eAuth.GET("/api/v1/product-category", controllers.GetCategoryControllers)
	eAuth.GET("/api/v1/product-category/:categoryId", controllers.DetailCategoryControllers)
	eAuth.PUT("/api/v1/product-category/:categoryId", controllers.EditCategoryControllers)
	eAuth.DELETE("/api/v1/product-category/:categoryId", controllers.DeleteCategoryControllers)

	/** FEATURE PRODUCT CATEGORY */
	eAuth.POST("/api/v1/product", controllers.CreateProductControllers)
	eAuth.GET("/api/v1/product", controllers.GetProductControllers)
	eAuth.GET("/api/v1/product/:categoryId", controllers.DetailProductControllers)
	eAuth.PUT("/api/v1/product/:categoryId", controllers.EditProductControllers)
	eAuth.DELETE("/api/v1/product/:categoryId", controllers.DeleteProductControllers)

	/* FEATURE CART */
	eAuth.POST("/api/v1/cart", controllers.CreateProductControllers)
	eAuth.DELETE("/api/v1/cart/:productId", controllers.DeleteProductControllers)

	/* FEATURE TRANSACTION */
	eAuth.POST("/api/v1/transaction", controllers.CreateProductControllers)
	eAuth.DELETE("/api/v1/transaction/:transactionId", controllers.DeleteProductControllers)

	return e
}
