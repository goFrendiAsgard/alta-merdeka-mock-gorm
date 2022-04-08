package controller

import (
	"fmt"
	"merdeka/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type DBEngine interface {
	Find(dest interface{}, conds ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
}

func NewPersonController(db DBEngine) PersonController {
	return PersonController{
		Db: db,
	}
}

type PersonController struct {
	Db DBEngine
}

func (pc PersonController) GetAllPerson(c echo.Context) error {
	persons := []model.Person{}
	tx := pc.Db.Find(&persons)
	if tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "cannot fetch persons",
		})
	}
	return c.JSON(http.StatusOK, persons)
}

func (pc PersonController) AddPerson(c echo.Context) error {
	person := model.Person{}
	c.Bind(&person)
	fmt.Println(person)
	tx := pc.Db.Save(&person)
	if tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "cannot add person",
		})
	}
	return c.JSON(http.StatusOK, person)
}
