package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pabloCode010/database_programming_project/models"
	"github.com/pabloCode010/database_programming_project/utils"
)

// Create a user in the database using the data sent in the request
func CreateUser(ctx echo.Context) error {
	var user models.User

	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	// Encrypt the password
	passwordEncrypted := utils.HashSHA256(*user.Password)
	user.Password = &passwordEncrypted

	// Execute the method to create the user
	_, err := user.Execute(1)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "Usuario creado correctamente",
		"user":    user,
	})
}
