package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}

type spanishBot struct{}

type chineseBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}
	ch := chineseBot{}
	printGreeting((eb))
	printGreeting(sb)
	printGreeting(ch)
}

func printGreeting(eb bot) {
	fmt.Println(eb.getGreeting())
}

func (englishBot) getGreeting() string {
	// very custom logic for generating an english greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func (chineseBot) getGreeting() string {
	return "hey chinese"
}
