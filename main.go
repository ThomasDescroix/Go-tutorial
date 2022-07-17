package main

import "fmt"

type Todo struct {
	Id        int
	Title     string
	Completed bool
}

type Toggleable interface {
	toggle()
}

func (t *Todo) toggle() {
	t.Completed = !t.Completed
}

func toggleable(t Toggleable) {
	t.toggle()
}

func main() {
	// map and struct
	a := map[string]string{"name": "John", "surname": "Doe"}
	a["name"] = "Jane"
	b := Todo{1, "Buy milk", false}
	b.toggle()
	toggleable(&b)
	fmt.Println(a, b)
}
