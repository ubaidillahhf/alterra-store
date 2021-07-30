package controllers

import (
	"alterra_store/lib/database"
	"alterra_store/models/carts"
	"alterra_store/validations"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateCartControllers(c echo.Context) error {

	var cartCreate carts.CartStruct
	c.Bind(&cartCreate)

	// Validasi Field
	errorValidate := validations.Validate(cartCreate)
	if errorValidate != nil {
		return errorValidate
	}

	createCartDB, err := database.CreateCart(cartCreate)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, BaseResponse(http.StatusCreated, "Add Cart Succes", createCartDB))
}

func DeleteCartControllers(c echo.Context) error {

	paramsCartId := c.Param("cartId")
	cartId, _ := strconv.Atoi(paramsCartId)
	_, e := database.DeleteCart(cartId)

	if e != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Get Data",
			nil,
		))
	}
	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Delete Cart",
		nil,
	))
}
