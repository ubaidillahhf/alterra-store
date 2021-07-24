package controllers

import (
	"alterra_store/configs"
	"alterra_store/helpers"
	"alterra_store/lib/database"
	"alterra_store/middlewares"
	"alterra_store/models/users"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func RegisterControllers(c echo.Context) error {

	var usersCreate users.UsersCreate
	c.Bind(&usersCreate)

	hash, _ := helpers.HashPassword(usersCreate.Password)

	var usersDB users.User
	usersDB.Name = usersCreate.Name
	usersDB.Address = usersCreate.Address
	usersDB.Email = usersCreate.Email
	usersDB.Password = hash

	err := configs.DB.Create(&usersDB).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, usersDB)
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
