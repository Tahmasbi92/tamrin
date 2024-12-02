package main

import "fmt"

func main() {

	cardsAce := []string{}

	cards := []string{"Ace Pik!", card()}

	// Adding to slices
	cardsAce = append(cards, "Ace khesht!!!")

	fmt.Println(cards)

	fmt.Println(cardsAce)
}

func card() string {
	return "Ace Dell!!"
}

/* Output:

[Ace Pik! Ace Dell!!]
[Ace Pik! Ace Dell!! Ace Khesht!!!]

*/
