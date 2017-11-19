package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
)

type CrudController interface {
	GetTable() *gorm.DB
	GetModel() interface{}
	GetModels() []interface{}
}

func CrudGetAll(c CrudController, w http.ResponseWriter) {
	models := c.GetModels()
	table := c.GetTable()
	table.Find(&models)
	res, err := json.Marshal(models)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(res)
}
