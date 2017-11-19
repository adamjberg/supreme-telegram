package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

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

func (c *OrganizationController) get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	model := models.Organization{
		ID: uint(id),
	}
	c.table.First(&model, model)
	res, err := json.Marshal(model)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write(res)
}

func (c *OrganizationController) getAll(w http.ResponseWriter, r *http.Request) {
	var models []models.Organization
	c.table.Find(&models)
	res, err := json.Marshal(models)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(res)
}

func (c *OrganizationController) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	model := models.Organization{
		ID: uint(id),
	}
	delete := c.table.Delete(&model)
	if delete.Error != nil {
		fmt.Println(delete.Error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func (c *OrganizationController) create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var model models.Organization
	err = json.Unmarshal(body, &model)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	create := c.table.Create(&model)
	if create.Error != nil {
		fmt.Println(create.Error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
