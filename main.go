package main

import "fmt"

type String string
type MyString String

func main() {
	var s MyString = "hello"
	ss := String(s)
	sss := string(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", ss)
	fmt.Printf("%T\n", sss)
}
