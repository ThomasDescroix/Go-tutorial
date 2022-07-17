package main

import "fmt"

func main() {
	// Pointer
	name := "John"
	name2 := &name
	*name2 = "Jane"
	fmt.Printf("Hello, %s %s\n", name, *name2)
}
