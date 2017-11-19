package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type OrganizationController struct {
}

func NewOrganizationController(router mux.Router) *OrganizationController {
	c := new(OrganizationController)

	s := router.PathPrefix("/organizations").Subrouter()

	s.HandleFunc("", c.getAll).Methods("GET")
	s.HandleFunc("/{id}", c.get).Methods("GET")
	s.HandleFunc("", c.create).Methods("POST")
	s.HandleFunc("/{id}", c.delete).Methods("DELETE")

	return c
}

func (c *OrganizationController) get(w http.ResponseWriter, r *http.Request) {

}

func (c *OrganizationController) getAll(w http.ResponseWriter, r *http.Request) {

}

func (c *OrganizationController) delete(w http.ResponseWriter, r *http.Request) {

}

func (c *OrganizationController) create(w http.ResponseWriter, r *http.Request) {

}
