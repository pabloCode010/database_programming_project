package showtimes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pabloCode010/database_programming_project/models"
)

// get all or filtered showtimes from the database
// If all the fields are empty, it will return all showtimes

func GetShowtimes(ctx echo.Context) error {
	var showTime models.ShowTime

	if err := ctx.Bind(&showTime); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	// Execute the method to get the showtimes
	showTimes, err := showTime.Execute(4)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message":   fmt.Sprintf("Se han encontrado %d funciones", len(showTimes)),
		"showtimes": showTimes,
	})
}
