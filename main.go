package main

import "fmt"

var k,v interface{}

func main() {
	k = 123
	v = "123"
	fmt.Printf("123\n");
	fmt.Print(k,v)
}
