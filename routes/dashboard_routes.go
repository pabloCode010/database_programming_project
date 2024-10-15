package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/pabloCode010/database_programming_project/controllers/dashboard"
)

func dashboardRoutes(e *echo.Echo) {
	e.GET("/dashboard/:entity", dashboard.RenderDashboard)
}
