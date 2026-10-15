package main

import "fmt"

// Declaring a new struct type
type Book struct {
	Title  string
	Author string
	Copies int
}

// func TestBookToString_FormatsBookInfoAsString() {
// 	input := Book{
// 		Title:  "Sea Room",
// 		Author: "Adam Nicolson",
// 		Copies: 2,
// 	}

// 	want := "Sea Room by Adam Nicolson - 2 copies"
// 	got := BookToString(input)

// 	if want != got {
// 		panic("BookToString: wrong result")
// 	}
// }

func BookToString(book Book) string {
	return fmt.Sprintf("%v b %v - %v copies", book.Title, book.Author, book.Copies)
}

func main() {
	fmt.Println("It's all good!")
}
