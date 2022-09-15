package pkg

import (
	"net/http"
	"simple-storage/api"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

// struct for session cookie
type Session struct {
	Username string
}

var sessions = map[string]Session{}

func DetermineContentType(data []string) bool {
	var found bool
	listContent := viper.GetStringMap("content")
	for _, v := range listContent {
		if data[0] == v {
			found = true
		}
	}

	return found
}

func Callback(code int, message string, data interface{}) (res api.Response, err error) {
	res.HttpCode = code
	res.Message = message
	res.Data = data

	return
}

func SetCookie(c echo.Context, name, path, value string) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Path = path
	cookie.Value = value
	c.SetCookie(cookie)
}

func GetCookie(c echo.Context, name string) (string, error) {
	cookie, err := c.Cookie(name)
	sessionToken := cookie.Value
	return sessionToken, err
}

func SetSessionName(sessionName, name string) {
	sessions[sessionName] = Session{
		Username: name,
	}
}

func GetSessionName(sessionName string) (string, bool) {
	userSession, exists := sessions[sessionName]
	return userSession.Username, exists
}

func DeleteSessionName(sessionName string) {
	delete(sessions, sessionName)
}
