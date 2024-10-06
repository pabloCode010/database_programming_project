package users

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pabloCode010/database_programming_project/models"
)

// Get all or filtered users from the database
// If all the fields are empty, it will return all users
func GetUSers(ctx echo.Context) error {
	var user models.User

	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	// Execute the method to get the users
	users, err := user.Execute(4)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": fmt.Sprintf("Se han encontrado %d usuarios", len(users)),
		"users":   users,
	})
}
