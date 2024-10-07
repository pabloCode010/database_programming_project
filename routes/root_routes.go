package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/pabloCode010/database_programming_project/controllers/root"
)

func rootRoutes(e *echo.Echo) {
	e.GET("/helloworld", root.HelloWorld)
}
