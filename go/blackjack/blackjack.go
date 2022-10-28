package blackjack

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
		case card == "ten" || card == "jack" || card == "queen" || card == "king":
			return 10
		default:
			return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	sumPlayerCards := ParseCard(card1) + ParseCard(card2)
	dealerCardValue := ParseCard(dealerCard)

	switch {
		case card1 == "ace" && card2 == "ace":
			return "P"
		case (sumPlayerCards == 21) && (dealerCard != "ace" && dealerCard != "ten" && dealerCard != "jack" && dealerCard != "queen" && dealerCard != "king"):
			return "W"
		case (sumPlayerCards == 21) && (dealerCard == "ace" || dealerCard == "ten" || dealerCard == "jack" || dealerCard == "queen" || dealerCard == "king"):
			return "S"
		case (sumPlayerCards >= 17 && sumPlayerCards <= 20):
			return "S"
		case (sumPlayerCards >= 12 && sumPlayerCards <= 16) && (dealerCardValue <= 6):
			return "S"
		case (sumPlayerCards >= 12 && sumPlayerCards <= 16) && (dealerCardValue >= 7):
			return "H"
		default:
			return "H"
	}
}
