package genres

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pabloCode010/database_programming_project/models"
)

// Get all or filtered genres from the database
// If name is empty, it will return all genres
func GetGenres(ctx echo.Context) error {
	var genre models.Genre

	if err := ctx.Bind(&genre); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	// Execute the method to get the genres
	genres, err := genre.Execute(4)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": fmt.Sprintf("Se han encontrado %d géneros", len(genres)),
		"genres":  genres,
	})

}
