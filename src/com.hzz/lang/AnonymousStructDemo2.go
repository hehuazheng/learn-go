package main

import "fmt"

type APerson struct {
	name string
	addr struct{
		province string
		code string
	}
}

type AddrNoname struct {
	province string
	code string
}
func main() {
	var a = APerson{"name1", AddrNoname{"hz", "bb"}}
	a.name = "name2"
	fmt.Println(a)
}
