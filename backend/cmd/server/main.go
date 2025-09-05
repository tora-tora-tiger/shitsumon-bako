package main

import (
	"net/http"
	
	"backend/pkg/schema"
	
	"github.com/labstack/echo/v4"
)

type ApiServer struct {}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}