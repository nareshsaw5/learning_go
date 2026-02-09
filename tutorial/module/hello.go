package main

import "fmt"

func sayHello(name string) (string, error) {
	message := fmt.Sprintf("Hello %v", name)
	return message, nil
}
