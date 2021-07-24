package controllers

import (
	"alterra_store/configs"
	"alterra_store/helpers"
	"alterra_store/lib/database"
	"alterra_store/middlewares"
	"alterra_store/models/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUserControllers(c echo.Context) error {

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
