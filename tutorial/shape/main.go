package main

import "fmt"

type Shape interface {
	getArea() float64
}
type Triangle struct {
	height float64
	base   float64
}

type Square struct {
	sideLength float64
}

func (s Square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t Triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

/*
*
* The receiver t *Triangle is a pointer here, hence it is pass by reference in Go
* Without *, it is just pass by value
 */
func (t *Triangle) doubleValue() {
	t.base = t.base * 2
	t.height = t.height * 2

}

func printArea(s Shape) {
	fmt.Println(s.getArea())
}

func main() {

	triangle := Triangle{
		height: 10,
		base:   5,
	}
	// triangle.doubleValue()
	printArea(triangle)

	sqaure := Square{
		sideLength: 10,
	}

	printArea(sqaure)

}
