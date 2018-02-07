package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// main function to boot up everything
func main() {

	CreateJSON()
	createRoutes()
}

// createRoutes - a function containing all mux logic
func createRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/", GetIndex).Methods("GET")
	router.HandleFunc("/scripts/{file}", GetScripts).Methods("GET")
	router.HandleFunc("/recache", CreateJSONCache).Methods("GET")
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")

	headersOk := handlers.AllowedHeaders([]string{"Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
