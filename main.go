package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sabarish-manoharan/demo/db"
	"github.com/spf13/viper"
)

func personCreate(w http.ResponseWriter, r *http.Request) {
	var p db.Person

	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := db.DB.Create(&p).Error; err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Added to the db")

}
func getPerson(w http.ResponseWriter, r *http.Request) {
	var persons []db.Person

	if err := db.DB.Find(&persons).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application-json")
	json.NewEncoder(w).Encode(persons)
}
func updatePerson(w http.ResponseWriter, r *http.Request) {
	var person db.Person
	vars := mux.Vars(r)
	id := vars["id"]
	if err := db.DB.First(&person, id).Error; err != nil {
		http.Error(w, "Person not found", http.StatusNotFound)
	}

	var updatePerson db.Person
	if err := json.NewDecoder(r.Body).Decode(&updatePerson); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	person.Name = updatePerson.Name
	person.Age = updatePerson.Age
	person.Occupation = updatePerson.Occupation

	if err := db.DB.Save(&person).Error; err != nil {
		http.Error(w, "Failed to update", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)

}
func deletePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idToRemove := vars["id"]

	if err := db.DB.Delete(&db.Person{}, idToRemove).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Employee Deleted Successfully"})

}
func main() {

	r := mux.NewRouter()
	c := cors.New(cors.Options{
		// ==AllowedOrigins: []string{"https://go-crud.netlify.app"},
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)
	port := getPort()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})
	r.HandleFunc("/person", personCreate).Methods("POST")
	r.HandleFunc("/persons", getPerson).Methods("GET")
	r.HandleFunc("/person/{id}", deletePerson).Methods("DELETE")
	r.HandleFunc("/person/{id}", updatePerson).Methods("PUT")
	http.Handle("/", handler)
	db.ConnectDB()
	fmt.Println("Server is connected in the port :  8000")
	http.ListenAndServe(":"+port, nil)
}

func getPort() string {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	port, ok := viper.Get("PORT").(string)
	if !ok {
		fmt.Println("Invalid type assertion")
	}
	return port
}
