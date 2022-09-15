package store

import (
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"simple-storage/pkg"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func Upload(c echo.Context, upload multipart.File, handler *multipart.FileHeader) (int, interface{}, error) {
	// check session cookie
	sessionToken, err := pkg.GetCookie(c, "session_token")
	if err != nil || sessionToken == "" {
		if err == http.ErrNoCookie || sessionToken == "" {
			return http.StatusUnauthorized, nil, errors.New("[Upload] - invalid credential")
		}
		return http.StatusBadRequest, nil, errors.New("[Upload] - bad request")
	}
	// check content type through helper
	checkType := pkg.DetermineContentType(handler.Header.Values("Content-Type"))
	if !checkType {
		return http.StatusBadRequest, nil, errors.New("[Upload] - invalid filetype")
	}
	// get file uploader name
	user, ok := pkg.GetSessionName(sessionToken)
	if !ok {
		return http.StatusUnauthorized, nil, errors.New("[Upload] - invalid credential")
	}
	// create file to local storage
	rawFilename := handler.Filename
	updFilename := strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(rawFilename, "_", "-"), " ", "-"))
	filename := user + "_" + "*" + "_" + updFilename
	folderName := strings.ReplaceAll(viper.GetString("storage.folder"), "/", "")
	createFile, err := os.CreateTemp(folderName, filename)
	if err != nil {
		return http.StatusServiceUnavailable, nil, errors.New("[Upload] - error create temporary file")
	}
	defer createFile.Close()
	// read bytes of file
	fileBytes, err := io.ReadAll(upload)
	if err != nil {
		return http.StatusServiceUnavailable, nil, errors.New("[Upload] - error read bytes of file")
	}
	createFile.Write(fileBytes)

	return http.StatusOK, nil, errors.New("[Upload] - upload success")
}

func List(c echo.Context) (int, interface{}, error) {
	// check session cookie
	sessionToken, err := pkg.GetCookie(c, "session_token")
	if err != nil || sessionToken == "" {
		if err == http.ErrNoCookie || sessionToken == "" {
			return http.StatusUnauthorized, nil, errors.New("[List] - invalid credential")
		}
		return http.StatusBadRequest, nil, errors.New("[List] - bad request")
	}
	_, ok := pkg.GetSessionName(sessionToken)
	if !ok {
		return http.StatusUnauthorized, nil, errors.New("[List] - invalid credential")
	}
	// get base directory
	dir, err := os.Getwd()
	if err != nil {
		return http.StatusServiceUnavailable, nil, errors.New("[List] - error get directory")
	}
	dir = dir + viper.GetString("storage.folder")
	// read director
	readDir, err := os.Open(dir)
	if err != nil {
		return http.StatusServiceUnavailable, nil, errors.New("[List] - error open directory")
	}
	// return files in directory
	allFiles, err := readDir.Readdir(0)
	if err != nil {
		return http.StatusServiceUnavailable, nil, errors.New("[List] - error read files in directory")
	}
	// iterate file with their info for the response
	var list ListFile
	var lists []ListFile
	baseUrl := c.Request().Host
	apiUri := viper.GetString("storage.uri")

	for _, file := range allFiles {
		filename := file.Name()
		splitFilename := strings.Split(filename, "_")
		uploaderName := splitFilename[0]
		filetime := file.ModTime()
		downloadUrl := baseUrl + apiUri + filename

		list.URL = downloadUrl
		list.UploadedAt = filetime
		list.Uploader = uploaderName
		lists = append(lists, list)
	}

	return http.StatusOK, lists, errors.New("[List] - success get list files")
}

func Download(c echo.Context, imageName string) (int, interface{}, error) {
	// get image name from param url
	splitName := strings.Split(imageName, "_")
	countName := len(splitName)
	downloadName := splitName[countName-1]
	// get base directory
	dir, err := os.Getwd()
	if err != nil {
		return http.StatusServiceUnavailable, nil, errors.New("[Download] - error get directory")
	}
	dir = dir + viper.GetString("storage.folder")
	fullPath := dir + imageName
	// check if file exist
	openFile, err := os.Open(fullPath)
	if err != nil {
		return http.StatusBadRequest, nil, errors.New("[Downlaod] - file not found")
	}
	defer openFile.Close()
	// download as attachment
	download := c.Attachment(fullPath, downloadName)

	return http.StatusOK, download, errors.New("[Download] - success access/download file")
}
