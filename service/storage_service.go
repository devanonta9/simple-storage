package service

import (
	"net/http"
	"simple-storage/internal/auth"
	"simple-storage/internal/store"
	"simple-storage/pkg"

	"github.com/labstack/echo/v4"
)

type SvcStorage struct{}

func (s SvcStorage) Login(c echo.Context) error {
	var creds auth.Credentials
	err := c.Bind(&creds)
	if err != nil {
		resp, _ := pkg.Callback(http.StatusUnauthorized, "[Login] - invalid credential", nil)
		return c.JSON(http.StatusUnauthorized, resp)
	}

	code, data, err := auth.Login(c, creds.Username, creds.Password)
	resp, _ := pkg.Callback(code, err.Error(), data)
	return c.JSON(code, resp)
}

func (s SvcStorage) Logout(c echo.Context) error {
	code, data, err := auth.Logout(c)
	resp, _ := pkg.Callback(code, err.Error(), data)
	return c.JSON(code, resp)
}

func (s SvcStorage) Upload(c echo.Context) error {
	upload, handler, err := c.Request().FormFile("file")
	if err != nil {
		resp, _ := pkg.Callback(http.StatusServiceUnavailable, "[Upload] - error retrieving the file", nil)
		return c.JSON(http.StatusServiceUnavailable, resp)
	}
	defer upload.Close()
	code, data, err := store.Upload(c, upload, handler)
	resp, _ := pkg.Callback(code, err.Error(), data)
	return c.JSON(code, resp)
}

func (s SvcStorage) List(c echo.Context) error {
	code, data, err := store.List(c)
	resp, _ := pkg.Callback(code, err.Error(), data)
	return c.JSON(code, resp)
}

func (s SvcStorage) Download(c echo.Context) error {
	key := c.Param("key")
	code, data, err := store.Download(c, key)
	resp, _ := pkg.Callback(code, err.Error(), data)
	return c.JSON(code, resp)
}
