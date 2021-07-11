package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()
	fmt.Println("---------------------------------")
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
	fmt.Println("------------------------------------")
	fmt.Println(cards.toString())

	cards.saveToFile("my_cards")

	retriveCards := newDeckFromFile("my_cards")
	retriveCards.print()
	fmt.Println("------------> shuffle")
	cards.shuffle()
	cards.print()
}
