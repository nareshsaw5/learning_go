package main

import (
	"fmt"

	"github.com/go_tutorial/greetings"
)

func main() {
	// Get a greeting message and print it.
	message, _ := greetings.Hello("Naresh")
	fmt.Println(message)

}
