package showtimes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pabloCode010/database_programming_project/models"
)

// Create a showtime in the database using the data sent in the request
func CreateShowtime(ctx echo.Context) error {
	var showTime models.ShowTime

	if err := ctx.Bind(&showTime); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	// Execute the method to create the showtime
	_, err := showTime.Execute(1)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, echo.Map{
		"message":  "Función creada exitosamente",
		"showtime": showTime,
	})
}
