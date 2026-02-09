package main

import (
	"fmt"

	"learning_go/greetings/service"
	"learning_go/greetings/utils"
	"learning_go/hello/calculator"
)

func main() {
	fmt.Println("Hello Naresh. This is simple example of module in the same project")
	fmt.Println(calculator.Add(2, 3))
	fmt.Println(service.SayHello("Naresh"))
	fmt.Println(utils.SayHello())
}
