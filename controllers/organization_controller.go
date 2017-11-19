package controllers

import (
	"net/http"

	"github.com/adamjberg/supreme-telegram/models"

	"github.com/jinzhu/gorm"

	"github.com/gorilla/mux"
)

type OrganizationController struct {
	table *gorm.DB
}

func NewOrganizationController(router *mux.Router, db *gorm.DB) *OrganizationController {
	c := new(OrganizationController)

	tableName := "organizations"
	c.table = db.Table(tableName)

	s := router.PathPrefix("/" + tableName).Subrouter()

	s.HandleFunc("", c.getAll).Methods("GET")
	s.HandleFunc("/{id}", c.get).Methods("GET")
	s.HandleFunc("", c.create).Methods("POST")
	s.HandleFunc("/{id}", c.delete).Methods("DELETE")

	return c
}

func (c *OrganizationController) GetTable() *gorm.DB {
	return c.table
}

func (c *OrganizationController) GetModel() models.Organization {
	return models.Organization{}
}

func (c *OrganizationController) GetModel() []models.Organization {
	return []models.Organization{}
}

func (c *OrganizationController) get(w http.ResponseWriter, r *http.Request) {

}

func (c *OrganizationController) getAll(w http.ResponseWriter, r *http.Request) {
	CrudGetAll(c, w)
}

func (c *OrganizationController) delete(w http.ResponseWriter, r *http.Request) {

}

func (c *OrganizationController) create(w http.ResponseWriter, r *http.Request) {

}
