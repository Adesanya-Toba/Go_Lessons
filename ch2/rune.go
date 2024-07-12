package main

import "fmt"

// The rune type is an alias for the int32 type, just as byte is an alias for uint8
func main() {
	var myfirstrune rune = 'e'
	var txt string = "hello there"
	fmt.Println(myfirstrune)
	fmt.Println(txt)
}
