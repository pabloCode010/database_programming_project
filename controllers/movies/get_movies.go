package movies

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pabloCode010/database_programming_project/models"
)

func GetMovies(ctx echo.Context) error {
	var movie models.Movie

	if err := ctx.Bind(&movie); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	movies, err := movie.Execute(4)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": fmt.Sprintf("Se han encontrado %d películas", len(movies)),
		"movies":  movies,
	})
}
