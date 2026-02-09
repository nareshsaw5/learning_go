package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("This is a demo of array")
	// create an array of int
	// var ar or ar :=(shortcut to create ar and assign value)
	var ar = []int{1, 2, 3, 4, 5}

	// print content of array
	Print(ar)

	naresh := "Naresh"
	fmt.Println(naresh)

	// OR
	var john = "John"
	fmt.Println(john)
	var charAr = strings.Split(john, "")
	for _, v := range charAr {
		fmt.Println(v)
	}
	stringArray := []string{"Naresh", "Mukesh"}
	PrintStringArray(stringArray)

	// create array by make function, it will have value initialize to 0, length and cap of 5
	ar_make := make([]int, 5)
	for _, v := range ar_make {
		fmt.Println(v)
	}
	// le of ar_make is 5
	fmt.Println(cap((ar_make)))

	// cap of ar_make is 5
	fmt.Println(cap((ar_make)))

	fmt.Println("========")
	// array copy/reference example
	a := [4]int{2, 4, 6, 8}
	b := a
	b[0] = 3 // this doesn't change a
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("===== slice ========")
	// slice copy/reference example, the x and y can have different len and cap but they will refer same underlying array
	x := []int{2, 4, 6, 8}
	y := x
	y[0] = 3 // this change x as both are referring same array
	fmt.Println(x)
	fmt.Println(y)
	//	x[6] = 10 // panic error as slice len is 4 but we are assigning value at 5th index, run time error
	fmt.Println("===== more for slice ========")
	slice := []int{1, 2, 3, 4, 5, 6}
	subSlice := slice[1:4] // sublice of 1st to 3rd index, exclude 4th index
	fmt.Println(subSlice)
	fmt.Println(len(subSlice), cap(subSlice))

}

func Print(ar []int) {
	for _, v := range ar {
		fmt.Println(v)
	}
}

func CopyArray(ar []int) {
	arCopy := make([]int, len(ar))
	copy(arCopy, ar)
}

func PrintStringArray(ar []string) {
	ar1 := strings.Join(ar, ",")
	fmt.Printf("ar: %v\n", ar)
	fmt.Printf("ar: %v\n", ar1)
}
