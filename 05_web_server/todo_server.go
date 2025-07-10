package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Todo struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

var todos []Todo

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	var t Todo
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	todos = append(todos, t)
	w.WriteHeader(http.StatusCreated)
}

func main() {
	http.HandleFunc("/todos", getTodos)
	http.HandleFunc("/add", addTodo)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
