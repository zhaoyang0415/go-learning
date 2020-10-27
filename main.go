package main

import "fmt"

func main() {
	v:=5
L1:
	goto L2

	fmt.Println(v)

	goto L1
	fmt.Println("==================")
L2:
		fmt.Println("hello world")
}
