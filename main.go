package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/ak1103dev/go-bank-api/docs"
	httpSwagger "github.com/swaggo/http-swagger"
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

// @title API Document
// @version 1.0
// @description Description
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /

func main() {
	fmt.Printf("listen on 8080")
	router = mux.NewRouter()
	initUserRoutes()
	initAccountRoutes()
	// Initialize db connection
	initDB()

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}
