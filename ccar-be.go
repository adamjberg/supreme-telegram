package main

import (
	"log"
	"net/http"

	controllers "github.com/adamjberg/supreme-telegram/controllers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	controllers.NewOrganizationController(router, db)
	log.Fatal(http.ListenAndServe(":3000", router))
}
