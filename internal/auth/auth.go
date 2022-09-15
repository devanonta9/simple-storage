package auth

import (
	"errors"
	"net/http"
	"simple-storage/pkg"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func Login(c echo.Context, username, password string) (int, interface{}, error) {

	// example with 3 users from config
	users := viper.GetStringMapString("credential")
	// check password based on username
	checkCreds, ok := users[username]
	if !ok || checkCreds != password || username == "" || password == "" {
		return http.StatusUnauthorized, nil, errors.New("[Login] - invalid credential")
	}
	sessionToken := uuid.New().String()

	// set session cookie
	pkg.SetCookie(c, "session_token", "/", sessionToken)
	pkg.SetSessionName(sessionToken, username)

	return http.StatusOK, nil, errors.New("[Login] - login success")
}

func Logout(c echo.Context) (int, interface{}, error) {
	// check session cookie
	sessionToken, err := pkg.GetCookie(c, "session_token")
	if err != nil || sessionToken == "" {
		if err == http.ErrNoCookie || sessionToken == "" {
			return http.StatusUnauthorized, nil, errors.New("[Logout] - invalid credential")
		}
		return http.StatusBadRequest, nil, errors.New("[Logout] - bad request")
	}
	_, ok := pkg.GetSessionName(sessionToken)
	if !ok {
		return http.StatusUnauthorized, nil, errors.New("[Upload] - invalid credential")
	}
	// delete session name
	pkg.DeleteSessionName(sessionToken)
	// set session cookie to empty
	pkg.SetCookie(c, "session_token", "/", "")

	return http.StatusOK, nil, errors.New("[Logout] - logout success")
}
