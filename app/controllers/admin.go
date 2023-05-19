package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	// Get user logic goes here
	return c.JSON(http.StatusOK, map[string]string{"name": "Admin"})
}
