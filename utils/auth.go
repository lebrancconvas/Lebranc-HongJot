package utils

import (
	"os"

	"github.com/labstack/echo/v4"
)

func Authenticated(username, password string, c echo.Context) (bool, error) {
	adminUsername := os.Getenv("API_USERNAME")
	adminPassword := os.Getenv("API_PASSWORD")

	if username == adminUsername && password == adminPassword {
		return true, nil
	}

	return false, nil
}
