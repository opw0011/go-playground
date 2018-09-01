package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Todo struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// in memory variables to act as db
var uuid uint = 1
var toDoStore []Todo = []Todo{}

func ListTodos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(toDoStore)
}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	rTodo := decodeJSONRequest(r)

	newTodo := Todo{ID: uuid, Title: rTodo.Title, Content: rTodo.Content}
	toDoStore = append(toDoStore, newTodo)
	uuid++
	json.NewEncoder(w).Encode(newTodo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	// Get id from url
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		panic(err)
	}

	rTodo := decodeJSONRequest(r)

	var updatedIndex int
	for i, todo := range toDoStore {
		if todo.ID == uint(id) {
			newTodo := Todo{ID: todo.ID, Title: rTodo.Title, Content: rTodo.Content}
			toDoStore = append(toDoStore[:i], toDoStore[i+1:]...)
			toDoStore = append(toDoStore, newTodo)
			updatedIndex = i
			break
		}
	}

	json.NewEncoder(w).Encode(toDoStore[updatedIndex])
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		panic(err)
	}

	var deletedTodo Todo
	for i, todo := range toDoStore {
		if todo.ID == uint(id) {
			toDoStore = append(toDoStore[:i], toDoStore[i+1:]...)
			deletedTodo = todo
			break
		}
	}

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
