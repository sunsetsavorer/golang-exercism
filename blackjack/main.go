package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {

	switch card {
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	case "ace":
		return 11
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {

	result := ""

	cardSum := ParseCard(card1) + ParseCard(card2)
	dealerCardPoints := ParseCard(dealerCard)

	hasPairOfAces := card1 == "ace" && card2 == "ace"
	hasBlackjack := cardSum == 21

	switch {
	case hasPairOfAces:
		result = "P"
	case hasBlackjack:
		if dealerCardPoints == 10 || dealerCardPoints == 11 {
			result = "S"
		} else {
			result = "W"
		}
	case cardSum >= 17 && cardSum <= 20:
		result = "S"
	case cardSum >= 12 && cardSum <= 16:
		if dealerCardPoints >= 7 {
			result = "H"
		} else {
			result = "S"
		}
	case cardSum <= 11:
		result = "H"
	}

	return result
}
