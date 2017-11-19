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

func get(c *OrganizationController, w http.ResponseWriter, r *http.Request) {

}

func getAll(c *OrganizationController, w http.ResponseWriter, r *http.Request) {

}

func delete(c *OrganizationController, w http.ResponseWriter, r *http.Request) {

}

func create(c *OrganizationController, w http.ResponseWriter, r *http.Request) {

}
