package main

import (
	"html/template"

	"github.com/labstack/echo/v4"
	"github.com/pabloCode010/database_programming_project/config"
	"github.com/pabloCode010/database_programming_project/database"
	"github.com/pabloCode010/database_programming_project/routes"
	"github.com/pabloCode010/database_programming_project/utils"
)

func init() {
	config.Config()
	database.Connect()
}

func main() {
	e := echo.New()

	// Template renderer
	e.Renderer = &utils.Template{
		Templates: template.Must(template.ParseGlob("views/**/*.html")),
	}

	// Static files
	e.Static("/public", "public")

	// Routes
	routes.EnableRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
