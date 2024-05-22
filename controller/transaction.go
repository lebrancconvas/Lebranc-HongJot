package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetExpenses(c echo.Context) error {
	return c.JSON(http.StatusOK, "Get Expenses")
}
