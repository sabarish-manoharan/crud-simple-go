package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Id   int    `json:"id"`
}

var persons []person

func personCreate(w http.ResponseWriter, r *http.Request) {
	var p person

	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(p)
	persons = append(persons, p)

}
func getPerson(w http.ResponseWriter, r *http.Request) {
	if len(persons) == 0 {
		fmt.Fprintln(w, "Data Not Found")
		return
	}

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
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintln(w, "conversion error")
		return
	}
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
	idToRemove, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintln(w, "String conversion error")
		return
	}
	found := false
	for i, p := range persons {
		if p.Id == idToRemove {
			persons = append(persons[:i], persons[i+1:]...)
			found = true
		}
	}
	if found {
		fmt.Fprintln(w, "Deleted!")
	} else {
		http.Error(w, "Element Not Found", http.StatusNotFound)
	}
}
func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})
	r.HandleFunc("/person", personCreate).Methods("POST")
	r.HandleFunc("/persons", getPerson).Methods("GET")
	r.HandleFunc("/person/{id}", deletePerson).Methods("DELETE")
	r.HandleFunc("/person/{id}", updatePerson).Methods("PUT")
	http.Handle("/", r)
	fmt.Println("Server is connected in the port 8000")
	http.ListenAndServe(":8000", nil)

}
