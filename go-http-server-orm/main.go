package main

import (
	"fmt"
	"log"
	"net/http"
	"playground/go-http-server-orm/db"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Index page")
}

// handle all the http requests
func handleRequests() {
	router := mux.NewRouter()

	router.HandleFunc("/", index).Methods("GET")

	// GET, POST, PUT, DELETE
	router.HandleFunc("/todos", ListTodos).Methods("GET")
	router.HandleFunc("/todos", AddTodo).Methods("POST")
	router.HandleFunc("/todos/{id:[0-9]+}", UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id:[0-9]+}", DeleteTodo).Methods("DELETE")

	log.Fatal(http.ListenAndServe("localhost:3000", router))
}

// initail db migration
func initMigration() {
	db.DB.AutoMigrate(&Todo{})
}

func main() {
	fmt.Println("The server starts")

	db.Open()
	defer db.Close()

	initMigration()

	handleRequests()
}
