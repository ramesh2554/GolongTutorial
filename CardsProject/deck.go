package main

/*

In  cards project we use receiver function , multiple return statements , convert string to byte of slice ,
save file to hard drive using io/ioutil package , retrieve file from hard drive
shuffle


 --------------------overview----------------

imported main package and some important packages like math , time , os ..
create a new custom type

newDeck() ----->  creates  a decks
print() ---->its a receiver function and used to print newDecks
deal() -----> uses a slices used to create a slice range using :
toString() -----> convert decks to strings
saveToFile() -----> we save into a file in local storage in string format
newDeckFromFile() ---> used to retrieve data from a local storage file and then we convert to string by default its come like slice of byte
*/
import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {

	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "club"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {

	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {

	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)

}
func newDeckFromFile(filename string) deck {

	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")

	return deck(s)
}

/*
generate random shuffle
*/
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i] // swap
	}
}
