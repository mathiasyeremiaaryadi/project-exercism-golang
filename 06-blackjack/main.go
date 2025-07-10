package main

import "fmt"

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch {
	case card == "ace":
		return 11
	case card == "two":
		return 2
	case card == "three":
		return 3
	case card == "four":
		return 4
	case card == "five":
		return 5
	case card == "six":
		return 6
	case card == "seven":
		return 7
	case card == "eight":
		return 8
	case card == "nine":
		return 9
	case (card == "ten" || card == "jack") || (card == "queen" || card == "king"):
		return 10
	}

	return 0
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	sumCard := ParseCard(card1) + ParseCard(card2)
	valueDealerCard := ParseCard(dealerCard)

	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case sumCard == 21:
		if dealerCard == "ace" || dealerCard == "queen" || dealerCard == "ten" {
			return "S"
		}
		return "W"
	case sumCard >= 17 && sumCard <= 20:
		return "S"
	case sumCard >= 12 && sumCard <= 16:
		if valueDealerCard >= 7 {
			return "H"
		}
		return "S"
	}

	return "H"
}

func main() {
	value := ParseCard("ace")
	fmt.Println(value)
	// Output: 11

	fmt.Println(FirstTurn("ace", "ace", "jack"))   // "P"
	fmt.Println(FirstTurn("ace", "king", "ace"))   // "S"
	fmt.Println(FirstTurn("five", "queen", "ace")) // "H"
}
