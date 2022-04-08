package main

import (
	"merdeka/controller"
	"merdeka/model"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		e.Logger.Fatal(err)
	}
	db.AutoMigrate(&model.Person{})

	// e.GET("/", controller.Hello)
	// //         package    function

	// personController := controller.NewPersonController(db)
	// //                  package    function
	// e.GET("/persons", personController.GetAllPerson)
	// e.POST("/persons", personController.AddPerson)

	controller.RegisterRoute(e, db)
	e.Logger.Fatal(e.Start(":3000"))
}
