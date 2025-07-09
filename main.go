package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type person struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Occupation string `json:"occupation"`
	Id         string `json:"id"`
}

var persons []person

func personCreate(w http.ResponseWriter, r *http.Request) {
	var p person

	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	persons = append(persons, p)
	json.NewEncoder(w).Encode(persons)

}
func getPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	json.NewEncoder(w).Encode(persons)
}
func updatePerson(w http.ResponseWriter, r *http.Request) {
	var p person

	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	found := false
	id := vars["id"]
	for i, person := range persons {
		if person.Id == id {
			persons[i] = p
			found = true
		}
	}
	if found {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(persons)
	} else {
		http.Error(w, "Element not found", http.StatusNotFound)
	}

}
func deletePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idToRemove := vars["id"]
	for i, p := range persons {
		if p.Id == idToRemove {
			persons = append(persons[:i], persons[i+1:]...)
		}
	}
}
func main() {

	r := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:5500"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})
	r.HandleFunc("/person", personCreate).Methods("POST")
	r.HandleFunc("/persons", getPerson).Methods("GET")
	r.HandleFunc("/person/{id}", deletePerson).Methods("DELETE")
	r.HandleFunc("/person/{id}", updatePerson).Methods("PUT")
	http.Handle("/", handler)
	fmt.Println("Server is connected in the port :  8000")
	http.ListenAndServe(":8000", nil)
}
