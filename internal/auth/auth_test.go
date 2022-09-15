package auth

import (
	"simple-storage/pkg"
	"testing"

	"github.com/labstack/echo/v4"
)

var c echo.Context

func TestLogin(t *testing.T) {
	username := "user"
	password := "pass"
	code, data, err := Login(c, username, password)
	t.Logf("Status Code : %d Data: %v Error Message: %s", code, data, err)
}

func TestLogout(t *testing.T) {
	pkg.SetCookie(c, "session_token", "/", "abcd")
	code, data, err := Logout(c)
	t.Logf("Status Code : %d Data: %v Error Message: %s", code, data, err)
}
