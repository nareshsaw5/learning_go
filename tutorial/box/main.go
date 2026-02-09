package main

import "fmt"

type Box interface {
	getArea() float64
}

type SquareBox struct {
	length  float64
	breadth float64
}

func (s SquareBox) getArea() float64 {
	return s.length * s.breadth
}

type CircularBox struct {
	radius float64
}

func (c CircularBox) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

func printArea(b Box) {
	fmt.Println(b.getArea())
}

func main() {

	circularBox := CircularBox{
		radius: 5,
	}
	printArea(circularBox)

	squareBox := SquareBox{
		length:  10,
		breadth: 20,
	}
	printArea(squareBox)

}
