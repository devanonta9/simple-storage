package store

import (
	"simple-storage/pkg"
	"testing"

	"github.com/labstack/echo/v4"
)

var c echo.Context

func TestUpload(t *testing.T) {
	pkg.SetCookie(c, "session_token", "/", "abcd")
	code, data, err := Upload(c, nil, nil)
	t.Logf("Status Code : %d Data: %v Error Message: %s", code, data, err)
}

func TestList(t *testing.T) {
	pkg.SetCookie(c, "session_token", "/", "abcd")
	code, data, err := List(c)
	t.Logf("Status Code : %d Data: %v Error Message: %s", code, data, err)
}

func TestDownload(t *testing.T) {
	code, data, err := Download(c, "image.jpg")
	t.Logf("Status Code : %d Data: %v Error Message: %s", code, data, err)
}
