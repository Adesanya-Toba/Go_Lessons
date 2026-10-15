package main

import "fmt"

// Declaring a new struct type
type Book struct {
	Title  string
	Author string
	Copies int
}

func BookToString(book Book) string {
	return fmt.Sprintf("%v by %v - %v copies", book.Title, book.Author, book.Copies)
}

func main() {
	book := Book{
		Title:  "Engineering in plain sight",
		Author: "Grady Hillhouse",
		Copies: 2,
	}
	fmt.Println(BookToString(book))
}
