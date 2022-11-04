package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'.
// Which is a slice of strings.
type deck []string

// Function newDeck of type deck defines suits and values and combines them together using a double for loop.
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// d deck = reciever, any variable of type "deck" now gets access to the "print" method. The actual copy of the deck we're working with is available in the function as a variable called 'd'.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// deal function take first agrument of type deck and secont argument of type int. Two vlaues will be returned both of type deck.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Takes the deck and turns it into a string.
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveTofile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 777)
}

// err is a value of type 'error', If nothing went wrong, it will have a value of 'nil'.
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s:= strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}