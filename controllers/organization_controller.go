package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type OrganizationController struct {
}

func NewOrganizationController(router *mux.Router) *OrganizationController {
	c := new(OrganizationController)

	s := router.PathPrefix("/organizations").Subrouter()

	s.HandleFunc("", c.getAll).Methods("GET")
	s.HandleFunc("/{id}", c.get).Methods("GET")
	s.HandleFunc("", c.create).Methods("POST")
	s.HandleFunc("/{id}", c.delete).Methods("DELETE")

	fmt.Println("Org controller")

	return c
}

func (c *OrganizationController) get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get")
	w.Write([]byte("Success"))
}

func (c *OrganizationController) getAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getAll")
	w.Write([]byte("Success"))
}

func (c *OrganizationController) delete(w http.ResponseWriter, r *http.Request) {

}

func (c *OrganizationController) create(w http.ResponseWriter, r *http.Request) {

}
