package main

// I is a shortcut to integer
type I int8

func (i *I) add(m I) {
	*i = *i + m
}
