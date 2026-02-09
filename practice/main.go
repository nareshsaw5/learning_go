package main

import (
	"fmt"

	"github.com/0101binarybard/practice/allocation"
	"github.com/0101binarybard/practice/maptest"
)

func main() {
	fmt.Println("My practice of Go")
	maptest.TestMap()
	allocation.NewTest()
	testFile := allocation.NewFile(20, "test")
	fmt.Println(testFile)
	fmt.Println("=======")
	ar := allocation.CreateSlice()
	fmt.Println(ar)
	fmt.Println("=======")
	array := [...]float64{7.0, 8.5, 9.1}
	x := allocation.Sum(&array) // Note the explit address-of operator
	fmt.Println(x)

	var sum float64
	for _, v := range array {
		sum += v
	}
	fmt.Println(sum)

}
