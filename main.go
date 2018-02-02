package main

import (
	"fmt"
	"github.com/daniel/rest-api-persons/db"
	"github.com/daniel/rest-api-persons/person"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	db.DBCon = db.InitDb("persosns.db")
	defer db.DBCon.Close()

	db.CreateTable(db.DBCon)
	router := httprouter.New()
	router.GET("/api/person/:id", person.FindPerson)
	router.POST("/api/person", person.CreatePerson)
	log.Fatal(http.ListenAndServe(":8080", router))
}
