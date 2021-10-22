package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	cards := newDeckFromFile("cards.txt")
	cards.shuffle()
	cards.print()
}

type deck []string

var cardSuits = []string{"Spades", "Diamonds", "Hearts", "Clubs"}
var cardValues = []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

func newDeck() deck {
	cards := deck{}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for idx, card := range d {
		fmt.Println(idx, card)
	}
}

func (d deck) saveToFile(filename string) error {
	deckString := d.toString()

	return ioutil.WriteFile(filename, []byte(deckString), 0666)
}

func deal(d deck, handSize int) (deck, deck) {
	dealed := d[:handSize]
	remaining := d[handSize:]

	return dealed, remaining
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func newDeckFromFile(fileName string) deck {
	byteSlice, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Err: ", err)
		os.Exit(1)
	}

	stringDeck := string(byteSlice)

	return deck(strings.Split(stringDeck, ","))
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)

	deckLen := len(d) - 1

	for i := range d {
		randIdx := randomizer.Intn(deckLen)

		d[i], d[randIdx] = d[randIdx], d[i]
	}
}
