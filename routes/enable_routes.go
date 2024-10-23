package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/pabloCode010/database_programming_project/middlewares"
)

func EnableRoutes(e *echo.Echo) {
	rootRoutes(e)
	dashboardRoutes(e)
	authRoutes(e)

	api := e.Group("/api/v1")

	// Api middlewares
	api.Use(middlewares.IsAuthenticated)
	api.Use(middlewares.IsAdmin)

	// Api routes
	usersRoutes(api)
	genresRoutes(api)
	moviesRoutes(api)
	showtimesRoutes(api)
}
