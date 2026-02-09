package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type called deck
// the deck is a slice of string
type deck []string

func createDeck() deck {
	var cards deck = deck{}
	var cardSuits = []string{"Spades", "Diamonds", "Hearts"}
	var cardValues = []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}

// create a function which prints slices in a deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	fmt.Println("deck size ", len(d))
	for i, card := range d {
		fmt.Println(i, card)
	}
	fmt.Println("============")
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(filename string) error {
	fmt.Println("saving to file ")
	// func WriteFile
	// func WriteFile(filename, data []byte, perm os.FilerMode) error
	// var values []string = d
	// d is slice of string, or array of string
	// loop over the deck/slice of string
	/**
	for _, card := range d {
		// each slice is a card
		fmt.Println(card)
		// covery a card - 'Spades of Two' into slice of byte b
		var b = []byte(card)
		// b will be something like [83 112 97 100 101 115 32 111 102 32 65 99 101]
		fmt.Println(b)
	}
	*/
	s := d.toString()
	// fmt.Print(s)
	var b = []byte(s)
	// fmt.Println(b)
	return ioutil.WriteFile(filename, b, 0666)
}

/**
* This returns all the cards of deck joined by separator ,
 */
func (d deck) toString() string {
	s := strings.Join(d, ", ")
	return s
}

func readDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error", err.Error())
		os.Exit(1) // exit the program, may not needed
	}
	// s := strings.Split(string(bs), ",")
	var s string = string(bs)
	var splittedString []string = strings.Split(s, ",")
	return deck(splittedString)

}

/**
* shuffle the cards in the deck
* change the card at position
 */
func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		// swap the value at two indexes
		d[i], d[newPosition] = d[newPosition], d[i]
	}

}
