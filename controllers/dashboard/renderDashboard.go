package dashboard

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RenderDashboard(c echo.Context) error {
	entity := c.Param("entity")

	return c.Render(http.StatusOK, "dashboard.html", map[string]interface{}{
		"Entity": entity,
	})
}
