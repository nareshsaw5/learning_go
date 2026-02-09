package main

import (
	"fmt"

	"github.com/go_tutorial/greetings"
)

func main() {

	naresh, _ := greetings.Hello("Naresh")
	fmt.Println(naresh)
}
