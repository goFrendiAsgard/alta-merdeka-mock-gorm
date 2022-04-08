package controller

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterRoute(e *echo.Echo, db *gorm.DB) {
	e.GET("/", Hello)

	personController := NewPersonController(db)
	e.GET("/persons", personController.GetAllPerson)
	e.POST("/persons", personController.AddPerson)
}
