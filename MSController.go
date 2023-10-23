package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func start(c echo.Context) error {
	return c.String(http.StatusOK, "Start")
}

func stop(c echo.Context) error {
	return c.String(http.StatusOK, "Stop")
}

func stats(c echo.Context) error {
	return c.String(http.StatusOK, "Stats")
}
