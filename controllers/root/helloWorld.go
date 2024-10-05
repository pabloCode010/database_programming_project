package root

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HelloWorld(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "Hello, World!",
	})
}
