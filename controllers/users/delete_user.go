package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pabloCode010/database_programming_project/models"
)

// Delete a user from the database using the param id
func DeleteUser(ctx echo.Context) error {
	var user models.User

	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	// Execute the method to delete the user
	_, err := user.Execute(2)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "Usuario eliminado correctamente",
	})
}
