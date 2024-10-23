package showtimes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pabloCode010/database_programming_project/models"
)

// delete a showtime from the database using the param id
func DeleteShowtime(ctx echo.Context) error {
	var showTime models.ShowTime

	if err := ctx.Bind(&showTime); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	// Execute the method to delete the showtime
	_, err := showTime.Execute(2)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "Función eliminada correctamente",
	})
}
