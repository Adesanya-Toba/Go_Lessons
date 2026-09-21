package main

import "fmt"

func printBook(title, author string, copies int) {
	fmt.Println(title, "by", author, "-", copies, "copies")

}

func main() {
	var title = "Killers of the flower moon"
	var author = "David Grann"
	var copies = 15
	printBook(title, author, copies)

	var x string
	fmt.Println(x)
}
