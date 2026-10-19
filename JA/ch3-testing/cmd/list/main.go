package main

import (
	"books"
	"fmt"
)

func main() {
	book := books.Book{
		Title:  "Engineering in plain sight",
		Author: "Grady Hillhouse",
		Copies: 2,
	}
	fmt.Println(books.BookToString(book))
}
