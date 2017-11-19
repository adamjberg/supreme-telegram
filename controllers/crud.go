package controllers

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

func GetAll(models []gorm.Model, table *gorm.DB) ([]byte, error) {
	table.Find(&models)
	res, err := json.Marshal(models)
	if err != nil {
		return nil, err
	}
	return res, nil
}
