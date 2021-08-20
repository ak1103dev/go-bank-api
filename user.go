package main

import (
	"encoding/json"
	"net/http"
	"strconv"
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

func create(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	db.Create(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
func find(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}
func findById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	userId := params["userId"]

	var user User
	db.First(&user, userId)
	json.NewEncoder(w).Encode(user)
}
func updateById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	userId := params["userId"]

	var user User
	db.First(&user, userId)

	var updatedUser User
	json.NewDecoder(r.Body).Decode(&updatedUser)
	db.Model(&user).Updates(updatedUser)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
func deleteById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userId"]
	id64, _ := strconv.ParseUint(userId, 10, 64)
	idToDelete := uint(id64)

	db.Where("user_id = ?", idToDelete).Delete(&User{})
	w.WriteHeader(http.StatusNoContent)
}

func initRoutes() {
	router.HandleFunc("/users", create).Methods("POST")
	router.HandleFunc("/users/{userId}", findById).Methods("GET")
	router.HandleFunc("/users", find).Methods("GET")
	router.HandleFunc("/users/{userId}", updateById).Methods("PUT")
	router.HandleFunc("/users/{userId}", deleteById).Methods("DELETE")
}
