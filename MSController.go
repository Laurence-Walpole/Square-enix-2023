package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func start(c echo.Context) error {
	job := createWorkerAndGetAJob("127.0.0.1")
	return c.String(http.StatusOK, "job id: "+strconv.FormatInt(job.ID, 10))
}

func pause(c echo.Context) error {
	return c.String(http.StatusOK, "Pause")
}

func stats(c echo.Context) error {
	return c.String(http.StatusOK, "Stats")
}
