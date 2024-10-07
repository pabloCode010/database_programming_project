package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/pabloCode010/database_programming_project/controllers"
	"github.com/pabloCode010/database_programming_project/controllers/auth"
)

func auth_routes(e *echo.Echo) {
	e.GET("/auth/sign-in", controllers.SendView("auth/sign-in.html"))
	e.POST("/auth/sign-in", auth.SignIn)
}
