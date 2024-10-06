package users

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/pabloCode010/database_programming_project/models"
	"github.com/pabloCode010/database_programming_project/utils"
)

// Update a user from the database using the param id
func UpdateUser(ctx echo.Context) error {
	var user models.User

	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	// Parse the new password query param
	newPassword, err := strconv.ParseBool(ctx.QueryParam("new_password"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	// If the new password is true, encrypt the new password
	if newPassword {
		passwordEncrypted := utils.HashSHA256(*user.Password)
		user.Password = &passwordEncrypted
	}

	// Execute the method to update the user
	_, err = user.Execute(3)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "Usuario actualizado correctamente",
		"user":    user,
	})

}
