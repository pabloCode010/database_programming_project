package showtimes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pabloCode010/database_programming_project/models"
)

// Update a showtime from the database using the param id
func UpdateShowtime(ctx echo.Context) error {
	var showTime models.ShowTime

	if err := ctx.Bind(&showTime); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	// Execute the method to update the showtime
	_, err := showTime.Execute(3)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message":  "Se ha actualizado la función",
		"showtime": showTime,
	})
}
