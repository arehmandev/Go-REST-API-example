package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// GetIndex - displays index page
func GetIndex(w http.ResponseWriter, r *http.Request) {
	indexpage, err := ioutil.ReadFile("index.html")
	if err != nil {
		log.Println(err)
		w.Write([]byte("Error - file not found"))
	}
	w.Write([]byte(indexpage))

}

// GetScripts - displays index page
func GetScripts(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	page, err := ioutil.ReadFile("scripts/" + params["file"])
	if err != nil {
		log.Println(err)
		w.Write([]byte("Error - file not found"))
	}
	w.Write([]byte(page))

}

// GetPeople - Display all from the people var
func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

// GetPerson - Display a single data
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

// CreatePerson - create a new item
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

// DeletePerson - deletes an item
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}

// CreateJSONCache - caches data into json file locally
func CreateJSONCache(w http.ResponseWriter, r *http.Request) {

	CreateJSON()
}
