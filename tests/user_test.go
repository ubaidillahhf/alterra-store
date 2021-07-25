package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var userJSON = `{"name":"","address":"sidoarjo","email":"jon@labstack.com"}`

func TestRegister(t *testing.T) {
	// Setup
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	// Assertions
	assert.Equal(t, http.StatusOK, rec.Code)
}
