package main

import (
	"log"
	"myBE/app/routes"
	"myBE/config"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.InitDB()
	e := echo.New()

	port := os.Getenv("APP_PORT")

	// Set up middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.ApiRoutes(e)
	e.Logger.Fatal(e.Start(":" + port))
}
