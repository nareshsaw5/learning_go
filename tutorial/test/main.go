package main

import "fmt"

func main() {

	var names = []string{"Naresh", "Suresh", "Dinesh"}

	for _, name := range names {
		fmt.Println(name)
	}

	var i int = 10
	fmt.Println(i)

	var ar = []int{1, 2, 2, 3, 3}
	for _, v := range ar {
		fmt.Println(v)
	}

}
