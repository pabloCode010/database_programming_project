package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/pabloCode010/database_programming_project/controllers"
	"github.com/pabloCode010/database_programming_project/controllers/dashboard"
	"github.com/pabloCode010/database_programming_project/middlewares"
)

func dashboardRoutes(e *echo.Echo) {
	dashboardPath := e.Group("/dashboard")

	// Middlewares
	dashboardPath.Use(middlewares.IsAuthenticated)
	dashboardPath.Use(middlewares.IsAdmin)

	// Routes
	dashboardPath.GET("", controllers.SendView("dashboard/menu.html"))
	dashboardPath.GET("/:entity", dashboard.RenderDashboard)
}
