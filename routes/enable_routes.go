package routes

import "github.com/labstack/echo/v4"

func EnableRoutes(e *echo.Echo) {
	root_routes(e)
	users_routes(e)
	auth_routes(e)
}
