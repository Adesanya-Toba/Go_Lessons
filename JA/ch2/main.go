package main

import "fmt"

func printBook(title, author string, copies int) {
	fmt.Println(title, "by", author, "-", copies, "copies")
}

func printBook2(book Book) {
	fmt.Printf("%v by %v - %v copies\n", book.Title, book.Author, book.Copies)
}

// Declaring a new struct type
type Book struct {
	Title  string
	Author string
	Copies int
}

func main() {
	var title = "Killers of the flower moon"
	var author = "David Grann"
	var copies = 15
	printBook(title, author, copies)

	var x string
	var y int = 90
	fmt.Println(x)
	fmt.Println(y)

	// Using structs
	var book Book

	book = Book{
		Title:  "Sea Room",
		Author: "Adam Nicolson",
		Copies: 2,
	}
	fmt.Println(book)

	var book2 = Book{
		Title:  "Hello world",
		Author: "Nice man",
		Copies: 21,
	}
	printBook2(book2)

	book3 := Book{
		Title:  "third book",
		Author: "Saul goodman",
		Copies: 0,
	}
	printBook(book3.Title, book3.Author, book3.Copies)
}
