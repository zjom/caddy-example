package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"key": "value"})
	})
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
