package main

import "fmt"

func main() {
	// Array
	a := []int{1, 2, 3}
	b := append(a, 7)
	b[1] = 4
	c := a[1:]
	a[1] = 6
	fmt.Println(a, b, c)
}
