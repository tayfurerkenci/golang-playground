package main

import "fmt"

type shape interface {
	getName() string
	getArea() float64
}

type triangle struct {
	name   string
	base   float64
	height float64
}

type square struct {
	name       string
	sideLength float64
}

func main() {
	t := triangle{name: "triangle", base: 10, height: 10}
	s := square{name: "square", sideLength: 10}

	printArea(t)
	printArea(s)
}

func (t triangle) getArea() float64 {
	return (t.base * t.height) / 2
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getName() string {
	return t.name
}

func (s square) getName() string {
	return s.name
}

func printArea(s shape) {
	fmt.Println(s.getName(), s.getArea())
}
