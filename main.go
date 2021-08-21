package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var router *mux.Router

func initDB() {
	var err error
	dsn := "host=localhost user=test password=test dbname=test port=9876 sslmode=disable TimeZone=Asia/Bangkok"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{}, &Account{})
}

func main() {
	router = mux.NewRouter()
	initUserRoutes()
	initAccountRoutes()
	// Initialize db connection
	initDB()

	log.Fatal(http.ListenAndServe(":8080", router))
}
