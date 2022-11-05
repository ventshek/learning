	// package main

	// import "fmt"

	// /* var = new variable. card = the name of the variable. string = only a string will ever be assigned to the variable. */
	// func main() {
	// 	// var card string = "Ace of Spades"
	// 	// := only used when defining a new variable
	// 	//card := "Ace of Spades"
	// 	//card := newCard()
	// 	//fmt.Println(card)
	// 	cards := []string{"Ace of Diamonds", newCard()}
	// 	cards = append(cards, "Six of Spades")
	// 	// i = index of this element in the array. card = Current card we're interating over. range cards = Take the slice of cards and loop over it. fmt.Println(card) Run this one time for each card in the slice.
	// 	for i, card := range cards {
	// 		fmt.Println(i, card)
	// 	}
	// 	//fmt.Println(cards)
	// }

	// // string = returns a value of type string when executing this function.
	// func newCard() string {
	// 	return "Five of Diamonds"
	// }

package main

//import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	//cards := newDeckFromFile("my_cards")
	//cards.print()
	//cards := newDeck()
	//cards.saveTofile("my_cards")
	//fmt.Println(cards.toString())
	//hand, remainingCards := deal(cards, 5)
	//hand.print()
	//remainingCards.print()
}
