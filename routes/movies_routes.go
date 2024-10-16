package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/pabloCode010/database_programming_project/controllers/movies"
)

func moviesRoutes(apiPath *echo.Group) {
	apiPath.POST("/movies", movies.CreateMovie)
	apiPath.DELETE("/movies/:id", movies.DeleteMovie)
	apiPath.PUT("/movies/:id", movies.UpdateMovie)
	apiPath.GET("/movies", movies.GetMovies)
}
