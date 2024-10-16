package movies

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pabloCode010/database_programming_project/models"
)

func UpdateMovie(ctx echo.Context) error {
	var movie models.Movie

	if err := ctx.Bind(&movie); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	_, err := movie.Execute(3)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "Película actualizada correctamente",
		"movie":   movie,
	})

}
