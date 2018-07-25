package main

import (
	"fmt"
)

type Shape interface {
	Draw()
}

type Rectangle struct {
}

func (this Rectangle) Draw() {
	fmt.Println("inside rectangle::draw()")
}

type Square struct {
}

func (this Square) Draw() {
	fmt.Println("inside square::draw()")
}

type Circle struct {
}

func (this Circle) Draw() {
	fmt.Println("inside circle::draw()")
}

type ShapeFactory struct {
}

func (this ShapeFactory) getShape(shapeType string) Shape {
	if shapeType == "c" {
		return Circle{}
	} else if shapeType == "r" {
		return Rectangle{}
	} else if shapeType == "s" {
		return Square{}
	}
	return nil
}
