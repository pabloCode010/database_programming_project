package middlewares

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/pabloCode010/database_programming_project/config"
)

func JwtMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		ContextKey:    "token",
		SigningMethod: "HS256",
		TokenLookup:   "cookie:Authorization",
		SigningKey:    []byte(config.JwtKey),
		Skipper: func(ctx echo.Context) bool {
			authorization, err := ctx.Cookie("Authorization")
			return err != nil || authorization.Value == ""
		},
	})
}
