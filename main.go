package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type User struct {
	// gorm.Model
	UserID    uint      `json:"userId" gorm:"primary_key"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func create(w http.ResponseWriter, r *http.Request) {}
func findById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	json.NewEncoder(w).Encode(users)
}
func find(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// params := mux.Vars(r)
	// userId := params["userId"]

	var user User
	json.NewEncoder(w).Encode(user)
}
func updateById(w http.ResponseWriter, r *http.Request) {}
func deleteById(w http.ResponseWriter, r *http.Request) {}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", create).Methods("POST")
	router.HandleFunc("/users/{userId}", findById).Methods("GET")
	router.HandleFunc("/users", find).Methods("GET")
	router.HandleFunc("/users/{userId}", updateById).Methods("PUT")
	router.HandleFunc("/users/{userId}", deleteById).Methods("DELETE")
	// Initialize db connection
	// initDB()

	log.Fatal(http.ListenAndServe(":8080", router))
}
