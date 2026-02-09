package main

import (
	"fmt"
)

func main() {
	hello, _ := sayHello("Naresh")
	fmt.Println(hello)
	hello_icm := hello.sayHello()

}
