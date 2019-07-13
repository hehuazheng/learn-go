package main

import "fmt"

type Shape struct {
	Name string
}

func (s Shape) show() {
	fmt.Println("my shape is " + s.getName())
}

func (s Shape) getName() string {
	return s.Name
}

type Square struct {
	Shape
	Length int
}

func (s Square) area() int {
	return s.Length * s.Length
}

func (s Square) getName() string {
	return "square"
}

func main() {
	var s Square = Square{Shape{Name: "shape"}, 3}

	s.show()
	fmt.Println(s.getName())
}
