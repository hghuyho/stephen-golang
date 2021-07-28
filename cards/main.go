package main

func main() {
	//cards := newDeck()
	// fmt.Println([]byte(cards.toString()))
	//cards.saveToFile("my_cards")
	cards := newDecFromFile("my_card")
	cards.print()
}
