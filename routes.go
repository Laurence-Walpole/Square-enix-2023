package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func routes(e *echo.Echo) {
	e.GET("/", root)
	e.GET("/start", start)
	e.GET("/pause", pause)
	e.GET("/stats", stats)
}

func root(c echo.Context) error {
	return c.String(http.StatusTeapot, "Undefined, try start, pause or stats.")
}
