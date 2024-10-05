package main

import (
	"github.com/labstack/echo/v4"
	"github.com/pabloCode010/database_programming_project/routes"
)

func main() {
	e := echo.New()

	// EnableRoutes(e)
	routes.EnableRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
