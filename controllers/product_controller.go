package controllers

import (
	"alterra_store/lib/database"
	"alterra_store/models/products"
	"alterra_store/validations"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

/* PRODUCT */
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

	return c.JSON(http.StatusCreated, BaseResponse(
		http.StatusCreated,
		"Success Create Product",
		createProductDB,
	))
}

func DetailProductControllers(c echo.Context) error {
	paramsProductId := c.Param("productId")
	productId, _ := strconv.Atoi(paramsProductId)

	categoryDB, e := database.GetProductDetail(productId)

	if e != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Get Data",
			nil,
		))
	}
	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Get Data by productId",
		categoryDB,
	))
}

func GetProductControllers(c echo.Context) error {

	var productData []products.Product
	var err error
	productData, err = database.GetProductAll()

	if err != nil {
		return c.JSON(http.StatusOK, BaseResponse(
			http.StatusInternalServerError,
			"Failed Get Data",
			productData,
		))
	}

	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Get Data Categories",
		productData,
	))
}

func EditProductControllers(c echo.Context) error {

	paramsProductId := c.Param("productId")
	categoryId, _ := strconv.Atoi(paramsProductId)
	var productEditDate products.ProductStruct
	c.Bind(&productEditDate)

	// Validasi Field
	errorValidate := validations.Validate(productEditDate)
	if errorValidate != nil {
		return errorValidate
	}

	userEdit, err := database.EditProduct(productEditDate, categoryId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, BaseResponse(http.StatusOK, "Sukses Edit Produk", userEdit))
}

func DeleteProductControllers(c echo.Context) error {

	paramsProductId := c.Param("productId")
	productId, _ := strconv.Atoi(paramsProductId)
	_, e := database.DeleteProduct(productId)

	if e != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Get Data",
			nil,
		))
	}
	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Delete Category",
		nil,
	))
}
