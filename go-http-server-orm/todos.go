package main

import (
	"encoding/json"
	"net/http"
	"playground/go-http-server-orm/db"
	"strconv"

	"github.com/gorilla/mux"
)

type Todo struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func ListTodos(w http.ResponseWriter, r *http.Request) {
	var todos []Todo
	db.DB.Find(&todos)
	json.NewEncoder(w).Encode(todos)
}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	rTodo := decodeJSONRequest(r)
	newTodo := Todo{Title: rTodo.Title, Content: rTodo.Content}
	db.DB.Create(&newTodo)
	json.NewEncoder(w).Encode(newTodo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		panic(err)
	}

	rTodo := decodeJSONRequest(r)

	var updatedTodo Todo
	db.DB.First(&updatedTodo, uint(id))

	if &updatedTodo != nil {
		updatedTodo.Content = rTodo.Content
		updatedTodo.Title = rTodo.Title
		db.DB.Save(&updatedTodo)
	}

	json.NewEncoder(w).Encode(updatedTodo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		panic(err)
	}

	var deletedTodo Todo
	db.DB.First(&deletedTodo, uint(id))
	db.DB.Delete(deletedTodo)

	json.NewEncoder(w).Encode(deletedTodo)
}

func decodeJSONRequest(r *http.Request) Todo {
	decoder := json.NewDecoder(r.Body)
	var rTodo Todo
	err := decoder.Decode(&rTodo)
	if err != nil {
		panic(err)
	}
	return rTodo
}
