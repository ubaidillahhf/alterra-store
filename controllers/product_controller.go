package controllers

import (
	"alterra_store/lib/database"
	"alterra_store/models/products"
	"alterra_store/validations"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateProductControllers(c echo.Context) error {

	var productCreate products.ProductStruct
	c.Bind(&productCreate)

	// Validasi Field
	errorValidate := validations.Validate(productCreate)
	if errorValidate != nil {
		return errorValidate
	}

	createProductDB, err := database.CreateProduct(productCreate)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, createProductDB)
}
