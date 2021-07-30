package controllers

import (
	"alterra_store/lib/database"
	"alterra_store/models/transactions"
	"alterra_store/validations"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateTransactionControllers(c echo.Context) error {
	var transactionCreate transactions.TransactionStruct
	c.Bind(&transactionCreate)

	// Validasi Field
	errorValidate := validations.Validate(transactionCreate)
	if errorValidate != nil {
		return errorValidate
	}

	createTransactionDB, err := database.CreateTransaction(transactionCreate)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, BaseResponse(http.StatusCreated, "Transaction Created", createTransactionDB))
}

func DetailTransactionControllers(c echo.Context) error {
	paramsTransactionId := c.Param("transactionId")
	transactionId, _ := strconv.Atoi(paramsTransactionId)

	categoryDB, e := database.GetCategoryDetail(transactionId)

	if e != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Get Data",
			nil,
		))
	}
	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Get Data by transactionId",
		categoryDB,
	))
}
