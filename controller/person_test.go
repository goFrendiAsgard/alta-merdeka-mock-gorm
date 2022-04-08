package controller

import (
	"bytes"
	"encoding/json"
	"merdeka/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewMockDB() *MockDB {
	return &MockDB{
		data: []interface{}{},
	}
}

type MockDB struct {
	data []interface{}
}

func (m *MockDB) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	return &gorm.DB{
		Error: nil,
	}
}

func (m *MockDB) Save(value interface{}) *gorm.DB {
	m.data = append(m.data, value)
	return &gorm.DB{
		Error: nil,
	}
}

func TestAddPerson(t *testing.T) {
	e := echo.New()

	m := NewMockDB()
	pc := NewPersonController(m)

	// add
	newPerson, _ := json.Marshal(map[string]string{
		"name":    "dono",
		"address": "depok",
	})
	addReq := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(newPerson))
	addReq.Header.Set("Content-Type", "application/json")
	addRec := httptest.NewRecorder()
	addContext := e.NewContext(addReq, addRec)
	addContext.SetPath("/")

	pc.AddPerson(addContext)

	var addedPerson model.Person
	if err := json.Unmarshal(addRec.Body.Bytes(), &addedPerson); err != nil {
		t.Errorf("unmarshalling returned person failed")
	}
	if addedPerson.Name != "dono" {
		t.Errorf("Expect: dono, Get: %s", addedPerson.Name)
	}
	if addedPerson.Address != "depok" {
		t.Errorf("Expect: depok, Get: %s", addedPerson.Address)
	}

	if len(m.data) != 1 {
		t.Errorf("Expect: 1, Get: %d", len(m.data))
	}

}
