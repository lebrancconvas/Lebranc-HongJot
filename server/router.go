package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lebrancconvas/Lebranc-HongJot/controller"
	"github.com/lebrancconvas/Lebranc-HongJot/utils"
)

func NewRouter() *echo.Echo {
	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	api := router.Group("/api")
	api.Use(middleware.BasicAuth(utils.Authenticated))
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/health", controller.Health)

			// Admin
			v1.GET("/expenses", controller.GetExpenses)
		}
	}

	return router
}
