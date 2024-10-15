package genres

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pabloCode010/database_programming_project/models"
)

// CreateGenre creates a genre in the database using the data sent in the request
func CreateGenre(ctx echo.Context) error {
	var genre models.Genre

	if err := ctx.Bind(&genre); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	// Execute the method to create the genre
	_, err := genre.Execute(1)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "Género creado correctamente",
		"genre":   genre,
	})
}
