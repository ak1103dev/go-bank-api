package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Account struct {
	// gorm.Model
	AccountID uint      `json:"accountId" gorm:"primary_key"`
	Name      string    `json:"name"`
	Owner     uint      `json:"owner" gorm:"foreignkey:UserID"`
	Balance   uint64    `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func accountCreate(w http.ResponseWriter, r *http.Request) {
	var account Account
	json.NewDecoder(r.Body).Decode(&account)
	db.Create(&account)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}
func accountFind(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var accounts []Account
	db.Preload("User").Find(&accounts)
	json.NewEncoder(w).Encode(accounts)
}
func accountFindById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	accountId := params["accountId"]

	var account Account
	db.First(&account, accountId)
	json.NewEncoder(w).Encode(account)
}
func accountUpdateById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	accountId := params["accountId"]

	var account Account
	db.First(&account, accountId)

	var updatedAccount Account
	json.NewDecoder(r.Body).Decode(&updatedAccount)
	db.Model(&account).Updates(updatedAccount)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}
func accountDeleteById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountId := params["accountId"]
	id64, _ := strconv.ParseUint(accountId, 10, 64)
	idToDelete := uint(id64)

	db.Where("account_id = ?", idToDelete).Delete(&Account{})
	w.WriteHeader(http.StatusNoContent)
}

func initAccountRoutes() {
	router.HandleFunc("/accounts", accountCreate).Methods("POST")
	router.HandleFunc("/accounts/{accountId}", accountFindById).Methods("GET")
	router.HandleFunc("/accounts", accountFind).Methods("GET")
	router.HandleFunc("/accounts/{accountId}", accountUpdateById).Methods("PUT")
	router.HandleFunc("/accounts/{accountId}", accountDeleteById).Methods("DELETE")
}
