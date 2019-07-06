package main

import "fmt"

type Book struct {
	name string
	author string
}

func main() {
	var b = Book{name:"book1"}
	var s string = tuple(b.author=="", "nil", "not nil").(string)
	fmt.Println("#", b.author, "#", s)
	fmt.Println(b)
}

func tuple(condition bool, a , b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}