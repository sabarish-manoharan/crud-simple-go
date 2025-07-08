package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

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
	if r.Method != "GET" {
		fmt.Fprintln(w, "Method Not Allowed")
		return
	}

	for _, p := range persons {
		fmt.Fprintf(w, "Id : %v \t Name : %v \t Age : %v\n", p.Id, p.Name, p.Age)
	}

}
func updatePerson(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		fmt.Fprintln(w, "Method not allowed")
		return
	}
	var p person

	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := p.Id

	for i, person := range persons {
		if person.Id == id {
			persons[i] = p
		}
	}
	for _, person := range persons {
		fmt.Fprintf(w, "Id : %v \t Name : %v \t Age : %v\n", person.Id, person.Name, person.Age)
	}

}
func deletePerson(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		fmt.Fprintln(w, "Method not allowed")
		return
	}

	vars := mux.Vars(r)
	clean := strings.Trim(vars["id"], "{}")
	idToRemove, err := strconv.Atoi(clean)
	if err != nil {
		fmt.Fprintln(w, "String conversion error")
		return
	}
	for i, p := range persons {
		if p.Id == idToRemove {
			persons = append(persons[:i], persons[i+1:]...)
		}
	}
	fmt.Fprintln(w, "Deleted!")
}
func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})
	r.HandleFunc("/person", personCreate)
	r.HandleFunc("/getPerson", getPerson)
	r.HandleFunc("/deletePerson/{id}", deletePerson)
	r.HandleFunc("/updatePerson", updatePerson)
	http.Handle("/", r)
	fmt.Println("Server is connected in the port 8000")
	http.ListenAndServe(":8000", nil)

}
