package controllers

import (
	"alterra_store/configs"
	"alterra_store/helpers"
	"alterra_store/models/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUsersControllers(c echo.Context) error {

	var usersCreate users.UsersCreate
	c.Bind(&usersCreate)

	hash, _ := helpers.HashPassword(usersCreate.Password)

	var usersDB users.Users
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
