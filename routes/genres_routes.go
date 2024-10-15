package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/pabloCode010/database_programming_project/controllers/genres"
)

func genresRoutes(apiPath *echo.Group) {
	apiPath.POST("/genres", genres.CreateGenre)
	apiPath.DELETE("/genres/:id", genres.DeleteGenre)
	apiPath.PUT("/genres/:id", genres.UpdateGenre)
	apiPath.GET("/genres", genres.GetGenres)
}
