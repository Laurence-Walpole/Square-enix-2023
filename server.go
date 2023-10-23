package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	//Create echo object
	e := echo.New()

	//Define routes
	routes(e)

	//Start server
	e.Logger.Fatal(e.Start(":1337"))
}
