package main

import (
	"myBE/app/routes"
	"myBE/config"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	config.InitDB()
	e := echo.New()

	port := os.Getenv("APP_PORT")

	// Set up middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.ApiRoutes(e)
	e.Logger.Fatal(e.Start(":" + port))
}
