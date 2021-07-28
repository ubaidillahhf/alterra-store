package controllers

import (
	"alterra_store/configs"
	"alterra_store/models/users"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitEcho() *echo.Echo {
	configs.InitDBTest()
	e := echo.New()
	return e
}

func CreateSeedUser() {
	var userDB users.User
	userDB.Name = "Test Name"
	userDB.Address = "Test Address"
	userDB.Email = "test@gmail.com"
	userDB.Password = "test"
	configs.DB.Create(&userDB)
}

func TestRegisterControllers(t *testing.T) {
	// Mock Data
	var userJSON = `{"name":"Test Name","address":"Test Address","email":"test@gmail.com","password":"123"}`
	var responseName = `Test Name`

	// Setup
	e := InitEcho()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/register")

	// Assertions
	if assert.NoError(t, RegisterControllers(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		body := rec.Body.String()
		baseResponse := users.UserResponseTest{}
		if err := json.Unmarshal([]byte(body), &baseResponse); err != nil {
			assert.Error(t, err, "Failed convert body to object")
		}

		// Minimal Equal Data Response
		assert.Equal(t, responseName, baseResponse.Data.Name)
	}
}

func TestLoginControllers(t *testing.T) {
	// Mock Data
	var userJSON = `{"email":"test@gmail.com","password":"123"}`
	var responseName = `Test Name`

	// Setup & Seed
	e := InitEcho()
	CreateSeedUser()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/login")

	// Assertions
	if assert.NoError(t, LoginControllers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		baseResponse := users.UserResponseTest{}
		if err := json.Unmarshal([]byte(body), &baseResponse); err != nil {
			assert.Error(t, err, "Failed convert body to object")
		}

		// Minimal Equal Data Response
		assert.Equal(t, responseName, baseResponse.Data.Name)
	}
}

func TestDetailUserControllers(t *testing.T) {

	// Setup & Seed
	e := InitEcho()
	CreateSeedUser()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/users/:userId")
	c.SetParamNames("userId")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, DetailUserControllers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		baseResponse := users.UserResponseTest{}
		if err := json.Unmarshal([]byte(body), &baseResponse); err != nil {
			assert.Error(t, err, "Failed convert body to object")
		}
	}
}
