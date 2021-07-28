package tests

import (
	"alterra_store/configs"
	"alterra_store/controllers"
	"alterra_store/models/users"
	"net/http"
	"net/http/httptest"
	"net/url"
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
	userDB.Email = "Email@gmail.com"
	userDB.Password = "password"
	configs.DB.Create(&userDB)
}

func TestRegister(t *testing.T) {
	e := InitEcho()

	apiUrl := "localhost:8000"
	resource := "/api/v1/users"
	data := url.Values{}
	data.Set("name", "Test Name")
	data.Set("address", "Test Address")
	data.Set("email", "Email@gmail.com")
	data.Set("password", "password")

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String() // "localhost:8000/api/v1/users"

	// requestByte, _ := json.Marshal(data.Encode())

	req := httptest.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode()))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/users")

	if assert.NoError(t, controllers.RegisterControllers(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		// body := rec.Body.String()
		// baseResponse := users.UserResponseTest{}
		// if err := json.Unmarshal([]byte(body), &baseResponse); err != nil {
		// 	assert.Error(t, err, "Failed convert body to object")
		// }
		// assert.Equal(t, http.StatusOK, baseResponse)
		// assert.Equal(t, 2, len(baseResponse.Data))
	}
}
