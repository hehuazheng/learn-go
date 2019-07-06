package main

import "fmt"
type Career struct {
	job string ""
}
type Person struct {
	Name string ""
	Addr string ""
	Career
}
func main() {
	var p = Person{"name", "addr", Career{"carrer2"}}
	fmt.Println(p)
}
