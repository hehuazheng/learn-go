package main

import "fmt"
type MyInt int

func (m *MyInt) Add(b MyInt) {
	var result = int(*m) + int(b)
	*m = MyInt(result)
}

func main() {
	var a MyInt = 1
	var b MyInt = 2
	a.Add(b)
	fmt.Println(a, b)
}
