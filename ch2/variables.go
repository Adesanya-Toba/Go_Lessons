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

	// Type conversions: Go enforces type casting, no implicit type conversion here!
	var a int = 10
	var b float64 = 90.25

	var sum1 float64 = float64(a) + b
	fmt.Printf("The sum is %.2f \n", sum1)

}
