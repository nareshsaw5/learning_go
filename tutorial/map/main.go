package main

import "fmt"

func main() {
	// var colors map[string]string
	// colors := make(map[string]string)
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"black": "#00000",
	}
	colors["white"] = "#ffffff"
	delete(colors, "black")
	print(colors)
}

func print(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, ":", value)
	}
}
