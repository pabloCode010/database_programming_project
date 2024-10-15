package routes

import "github.com/labstack/echo/v4"

func EnableRoutes(e *echo.Echo) {
	rootRoutes(e)
	dashboardRoutes(e)
	authRoutes(e)

	api := e.Group("/api/v1")
	usersRoutes(api)
	genresRoutes(api)
}
