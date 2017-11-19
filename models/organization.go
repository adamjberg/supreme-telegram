package models

import "github.com/jinzhu/gorm"

type Organization struct {
	gorm.Model
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
