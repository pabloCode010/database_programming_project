package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/pabloCode010/database_programming_project/controllers/users"
)

// users_routes defines the routes for the users
func usersRoutes(apiPath *echo.Group) {
	apiPath.POST("/users", users.CreateUser)       // 1: Create User
	apiPath.DELETE("/users/:id", users.DeleteUser) // 2: Delete User
	apiPath.PUT("/users/:id", users.UpdateUser)    // 3: Update User
	apiPath.GET("/users", users.GetUSers)          // 4: Get Users or Filter
}
