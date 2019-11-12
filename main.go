package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func connectDb() {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect")
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/", helloWorld).Methods("GET")
	router.HandleFunc("/users", AllUsers).Methods("GET")
	router.HandleFunc("/user/{name}/{email}", NewUser).Methods("POST")
	router.HandleFunc("/user/{name}", DeleteUser).Methods("DELETE")
	router.HandleFunc("/user/{name}/{email}", UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	fmt.Println("Go ORM example listening on :8081")

	connectDb()
	defer db.Close()

	db.AutoMigrate(&User{})

	handleRequests()
}
