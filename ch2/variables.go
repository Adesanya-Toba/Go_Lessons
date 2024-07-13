package main

import "fmt"

func main() {
	// Creating new variables
	var x int = 20
	y := 21

	// Declaring variables
	x *= 2 // = 40
	fmt.Println(x, y)

	// Using a function from a different package. Neat!
	rune_func()
}
