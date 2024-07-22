package main

import (
	"fmt"
	"slices"
)

func main() {
	/********************* ARRAYS ***************************/
	// Declare an empty array
	var x1 [3]int // 3 element array with zero default values

	// Initializing on declaration
	var x2 = [3]int{1, 2}

	// Initializing a sparse array, with most elements being zero
	var x3 = [...]int{1, 2, 4: 8, 78, 9: 10} // where x:y, is x = index, y = value

	// Another to declare an array of indefinite size
	var x4 = [...]int{1, 2, 3}

	// Array equality
	fmt.Println(x2 == x4) // Prints true

	fmt.Println(x1, x2, x3, x4)
	fmt.Println(len(x3))

	// fmt.Println(x4[80]) // Raises invalid argument, out of bounds error.

	// var y int = 3
	// var x5 = [y]int{2, 2} // Go sees the size of the array as part of the type.

	/********************* SLICES ***************************/
	var y1 = []int{1, 2, 4} // Creates a slice of three ints using a slice literal.
	fmt.Println(y1)

	// We can also declare a slice without using a literal
	var y2 []int // assigns y2 a nil value

	// The slice type is not comparable, the only thing you can compare it to is nil
	fmt.Println(y2 == nil)

	x := []int{1, 2, 3, 4, 5}
	y := []int{1, 2, 3, 4, 5}
	z := []int{1, 2, 3, 4, 5, 6}
	s := []string{"a", "b", "c"}
	fmt.Println(slices.Equal(x, y)) // prints true
	fmt.Println(slices.Equal(x, z)) // prints false
	fmt.Println(slices.Equal(x, s)) // does not compile
}
