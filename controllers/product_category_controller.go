package controllers

import (
	"alterra_store/lib/database"
	"alterra_store/models/productCategories"
	"alterra_store/validations"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

/** PRODUCT CATEGORIES */
func CreateCategoryControllers(c echo.Context) error {

	var categoryCreate productCategories.ProductCategoryStruct
	c.Bind(&categoryCreate)

	// Validasi Field
	errorValidate := validations.Validate(categoryCreate)
	if errorValidate != nil {
		return errorValidate
	}

	createCategoryDB, err := database.CreateCategory(categoryCreate)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, BaseResponse(
		http.StatusCreated,
		"Success Create Category",
		createCategoryDB,
	))
}

func DetailCategoryControllers(c echo.Context) error {
	paramsCategoryId := c.Param("categoryId")
	categoryId, _ := strconv.Atoi(paramsCategoryId)

	categoryDB, e := database.GetCategoryDetail(categoryId)

	if e != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Get Data",
			nil,
		))
	}
	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Get Data by categoryId",
		categoryDB,
	))
}

func GetCategoryControllers(c echo.Context) error {

	var categoryData []productCategories.ProductCategory
	var err error
	categoryData, err = database.GetCategoryAll()

	if err != nil {
		return c.JSON(http.StatusOK, BaseResponse(
			http.StatusInternalServerError,
			"Failed Get Data",
			categoryData,
		))
	}

	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Get Data Categories",
		categoryData,
	))
}

func EditCategoryControllers(c echo.Context) error {

	paramsCategoryId := c.Param("categoryId")
	categoryId, _ := strconv.Atoi(paramsCategoryId)
	var categoryEditData productCategories.ProductCategoryStruct
	c.Bind(&categoryEditData)

	// Validasi Field
	errorValidate := validations.Validate(categoryEditData)
	if errorValidate != nil {
		return errorValidate
	}

	categoryEdit, err := database.EditCategory(categoryEditData, categoryId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Edit Category",
		categoryEdit,
	))
}

func DeleteCategoryControllers(c echo.Context) error {

	paramsCategoryId := c.Param("categoryId")
	categoryId, _ := strconv.Atoi(paramsCategoryId)
	_, e := database.DeleteCategory(categoryId)

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
