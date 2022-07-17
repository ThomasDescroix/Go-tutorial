package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	Id        int
	Title     string
	Completed bool
}

func main() {
	// Chanel
	ch := make(chan string)
	errCh := make(chan error)
	go func() {
		title, err := fetchTodoTitle()
		if err != nil {
			errCh <- err
			return
		}
		ch <- title
		return
	}()
	fmt.Println("Hello, World!")
	defer close(ch)
	defer close(errCh)
	select {
	case err := <-errCh:
		panic(err)
	case title := <-ch:
		fmt.Println(title)
	}
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
