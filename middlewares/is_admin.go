package middlewares

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		token := ctx.Get("token").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)

		if claims["role"] != "ADMINISTRADOR" {
			return ctx.Redirect(http.StatusFound, "/home")
		}

		return next(ctx)
	}
}
