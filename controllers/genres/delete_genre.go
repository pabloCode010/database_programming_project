package genres

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pabloCode010/database_programming_project/models"
)

// DeleteGenre deletes a genre from the database using the param id
func DeleteGenre(ctx echo.Context) error {
	var genre models.Genre

	if err := ctx.Bind(&genre); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	// Execute the method to delete the genre
	_, err := genre.Execute(2)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "Género eliminado correctamente",
	})
}
