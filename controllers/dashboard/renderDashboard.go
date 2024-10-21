package dashboard

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var dashboardTitles = map[string]string{
	"users":  "Panel de Usuarios",
	"movies": "Panel de Películas",
	"genres": "Panel de Géneros(Películas)",
}

func RenderDashboard(c echo.Context) error {
	entity := c.Param("entity")

	return c.Render(http.StatusOK, "dashboard.html", map[string]interface{}{
		"Entity": entity,
		"Title":  dashboardTitles[entity],
	})
}
