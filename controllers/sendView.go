package controllers

import (
	"path"

	"github.com/labstack/echo/v4"
)

func SendView(file string) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.File(path.Join("views", file))
	}
}
