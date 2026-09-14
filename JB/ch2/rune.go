package main

import "fmt"

// The rune type is an alias for the int32 type, just as byte is an alias for uint8
func rune_func() {
	var myfirstrune rune = 'e'
	var mysecondrune int32 = 'e' // same as above just a bit confusing
	var txt string = "hello there"
	fmt.Println(myfirstrune)
	fmt.Println(mysecondrune)
	fmt.Println(txt)
}
