package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

type Todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	// Web Server
	color.Cyan("server started on port 8080")
	http.HandleFunc("/", homeHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	title, err := fetchTodoTitle()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error: %s", err)))
		return
	}
	w.Write([]byte(title))
}

func fetchTodoTitle() (string, error) {
	r, err := http.Get("https://jsonplaceholder.typicode.com/todos?_limit=5")
	if err != nil {
		return "", fmt.Errorf("error fetching todos: %s", err)
	}

	defer r.Body.Close()
	var todos []Todo
	err = json.NewDecoder(r.Body).Decode(&todos)
	if err != nil {
		return "", fmt.Errorf("error decoding todos: %s", err)
	}
	if len(todos) > 0 {
		return todos[0].Title, nil
	}
	return "", fmt.Errorf("no todos found")
}
