package main

import "fmt"

func main() {
	// slice of type string. Slice are Array in go
	// cards := deck{getCards(), "Test 1", "Test2"}
	var cards deck = createDeck()
	cards.print()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// fmt.Println("============")
	// remainingCards.print()
	// hand.saveToFile("hand.txt")
	// remainingCards.saveToFile("remainingCards.txt")
	fmt.Println("============")
	// var d deck = readDeckFromFile("hand.txt")
	// d.print()

	// d = readDeckFromFile("remainingCards.txt")
	cards.shuffle()
	fmt.Println("After shuffle")
	cards.print()
}

func getCards() string {
	return "Hello Cards"
}
