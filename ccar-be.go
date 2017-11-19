package main

import (
	"log"
	"net/http"

	controllers "github.com/adamjberg/supreme-telegram/controllers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	controllers.NewOrganizationController(router)
	log.Fatal(http.ListenAndServe(":3000", router))
}
