package movies

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pabloCode010/database_programming_project/models"
)

func CreateMovie(ctx echo.Context) error {
	var movie models.Movie

	if err := ctx.Bind(&movie); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	_, err := movie.Execute(1)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, echo.Map{
		"message": "Película creada exitosamente",
		"movie":   movie,
	})
}
