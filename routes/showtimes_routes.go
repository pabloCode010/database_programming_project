package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/pabloCode010/database_programming_project/controllers/showtimes"
)

func showtimesRoutes(apiPath *echo.Group) {
	apiPath.POST("/showtimes", showtimes.CreateShowtime)
	apiPath.DELETE("/showtimes/:id", showtimes.DeleteShowtime)
	apiPath.PUT("/showtimes/:id", showtimes.UpdateShowtime)
	apiPath.GET("/showtimes", showtimes.GetShowtimes)
}
