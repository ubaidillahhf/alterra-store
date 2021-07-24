package controllers

import (
	"alterra_store/lib/database"
	"alterra_store/middlewares"
	"alterra_store/models/users"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

func RegisterControllers(c echo.Context) error {
	v := validator.New()
	var usersCreate users.UserCreate
	c.Bind(&usersCreate)
	if errValidation := v.Struct(usersCreate); errValidation != nil {
		return errValidation
	}

	userDB, err := database.RegisterUser(usersCreate)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, userDB)
}

func LoginController(c echo.Context) error {

	userLogin := users.UserLogin{}
	c.Bind(&userLogin)

	userDB, e := database.LoginUser(userLogin)
	if e != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Get Data",
			nil,
		))
	}

	token, _ := middlewares.GenerateTokenJWT(userDB.Id)

	var userResponse = users.UserResponse{
		Id:        userDB.Id,
		Name:      userDB.Name,
		Email:     userDB.Name,
		CreatedAt: userDB.CreatedAt,
		UpdatedAt: userDB.UpdatedAt,
		DeletedAt: userDB.DeletedAt,
		Token:     token,
	}
	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Get Data",
		userResponse,
	))

}

func DetailUserControllers(c echo.Context) error {
	userId := middlewares.GetUserIdFromJWT(c)

	userDB, e := database.GetUserDetail(userId)
	paramsUserId := c.Param("userId")
	convertToInt, _ := strconv.Atoi(paramsUserId)

	if convertToInt != userDB.Id {
		return c.JSON(http.StatusBadRequest, BaseResponse(
			http.StatusBadRequest,
			"Bad Request Url Params",
			nil,
		))
	}
	if e != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Get Data",
			nil,
		))
	}
	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Get Data UserId",
		userDB,
	))
}

func GetUserControllers(c echo.Context) error {

	var userData []users.User
	var err error
	userData, err = database.GetDataUserAll()

	if err != nil {
		return c.JSON(http.StatusOK, BaseResponse(
			http.StatusInternalServerError,
			"Failed Get Data",
			userData,
		))
	}

	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Get Data News",
		userData,
	))
}

func EditUserControllers(c echo.Context) error {

	userId := middlewares.GetUserIdFromJWT(c)
	var userEditData users.UserEdit
	c.Bind(&userEditData)

	confirmedUser, _ := database.CheckHashPassword(userEditData.ConfirmPassword, userId)
	if !confirmedUser {
		return c.JSON(http.StatusUnauthorized, "Password Konfirmasi Salah")
	}

	userEdit, err := database.EditUser(userEditData, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, userEdit)
}
