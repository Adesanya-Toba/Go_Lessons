package main

import "fmt"

func main() {
	// Creating new variables
	var x int = 20
	var x1 = 10 // variable type is inferred from assignment
	var x2 int  // declaring and assigning a variable its zero value
	fmt.Println(x1, x2)

	var y1, y2 = 3, 10       // assigning multiple variables
	var y3, y4 int           // declaring and assigning zero value to multiple variables
	var y5, y6 = 56, "hello" // assigning variables of different types
	fmt.Println(y1, y2, y3, y4)
	fmt.Println(y5, y6)

	// Declare multiple variables at once using a Declaration list
	var (
		z1     int
		z2         = 30
		z3     int = 60
		z4, z5     = 35, "myagi"
		z6, z7 string
	)
	fmt.Println(z1, z2, z3, z4, z5, z6, z7)

	// When inside a function, use the := operator to replace a var
	// declaration that uses type inference
	y := 21
	// var y int = 21 // same as above

	a1, a2 := 79, "hello james"
	fmt.Println(a1, a2)

	// The := operator allows for one trick var does not
	// It allows you to assign values to existing variables too. As long as
	// at least one new variable is on the left-hand side of the :=, any of the other variables can already exist.
	a1, a3 := 90, 78

	// You can also just assign a new with = anyway
	a1 = 80
	fmt.Println(a1, a3)

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
