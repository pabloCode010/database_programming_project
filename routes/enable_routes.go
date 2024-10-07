package routes

import "github.com/labstack/echo/v4"

func EnableRoutes(e *echo.Echo) {
	rootRoutes(e)
	usersRoutes(e)
	authRoutes(e)
}
