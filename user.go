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

// userCreate godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body User true "Create user"
// @Success 200 {object} User
// @Router /users [post]
func userCreate(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	db.Create(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// userFind godoc
// @Summary Get details of all users
// @Description Get details of all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} User
// @Router /users [get]
func userFind(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func userFindById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	userId := params["userId"]

	var user User
	db.First(&user, userId)
	json.NewEncoder(w).Encode(user)
}

// userUpdateById godoc
// @Summary Update user identified by the given userId
// @Description Update the user corresponding to the input userId
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body User true "Update user"
// @Param userId path int true "ID of the user to be updated"
// @Success 200 {object} User
// @Router /users/{userId} [put]
func userUpdateById(w http.ResponseWriter, r *http.Request) {
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

// userDeleteById godoc
// @Summary Delete user identified by the given userId
// @Description Delete the user corresponding to the input userId
// @Tags users
// @Accept  json
// @Produce  json
// @Param userId path int true "ID of the user to be deleted"
// @Success 204 "No Content"
// @Router /users/{userId} [delete]
func userDeleteById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userId"]
	id64, _ := strconv.ParseUint(userId, 10, 64)
	idToDelete := uint(id64)

	db.Where("user_id = ?", idToDelete).Delete(&User{})
	w.WriteHeader(http.StatusNoContent)
}

func initUserRoutes() {
	router.HandleFunc("/users", userCreate).Methods("POST")
	router.HandleFunc("/users/{userId}", userFindById).Methods("GET")
	router.HandleFunc("/users", userFind).Methods("GET")
	router.HandleFunc("/users/{userId}", userUpdateById).Methods("PUT")
	router.HandleFunc("/users/{userId}", userDeleteById).Methods("DELETE")
}
