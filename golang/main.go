package main

import (
	"net/http"
	"os"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "work")
	})

	e.GET("/value1", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "return: value_01"})
	})

	e.GET("/value2", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "return: value_02"})
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "7008"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}