package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		token := ctx.Get("token")

		if token == nil {
			return ctx.Redirect(http.StatusFound, "/auth/sign-in")
		}

		return next(ctx)
	}
}
