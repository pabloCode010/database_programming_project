package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pabloCode010/database_programming_project/jwt"
	"github.com/pabloCode010/database_programming_project/models"
	"github.com/pabloCode010/database_programming_project/utils"
)

func SignIn(ctx echo.Context) error {
	var user models.User

	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	encryptedPassword := utils.HashSHA256(*user.Password)
	user.Password = &encryptedPassword

	_, err := user.Execute(5)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{
			"error": err.Error(),
		})
	}

	token, err := jwt.New(user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	ctx.SetCookie(&http.Cookie{
		Name:     "Authorization",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
	})

	return ctx.JSON(http.StatusOK, echo.Map{
		"message":  "Autenticación exitosa",
		"redirect": "/home",
	})
}
