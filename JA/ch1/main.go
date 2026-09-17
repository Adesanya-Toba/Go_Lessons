package main

import "fmt"

func hello() {
	fmt.Println("Hello, world!")
}

func main() {
	fmt.Println("Books in stock:")
	var title = "Alice in Wonderland"
	var author = "Kimmy K3"
	var book = "'Master and Commander', by Patrick O'Brian"
	fmt.Println(book)
	fmt.Println(title, "by", author)
	book = "'A Morbid Taste for Bones', by Ellis Peters"
	title = "A new title"
	author = "A new author"
	fmt.Println(book)
	fmt.Println(title, "by", author)
}
