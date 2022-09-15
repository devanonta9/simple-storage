package router

import (
	service "simple-storage/service"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) (*echo.Echo, error) {

	svc := service.SvcStorage{}
	authGroup := e.Group("auth")
	authGroup.POST("/login", svc.Login)
	authGroup.GET("/logout", svc.Logout)

	storeGroup := e.Group("store")
	storeGroup.POST("/upload", svc.Upload)
	storeGroup.GET("/list", svc.List)
	storeGroup.GET("/download/:key", svc.Download)

	return e, nil
}
