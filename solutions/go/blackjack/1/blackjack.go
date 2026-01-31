package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	res := 0
    
	switch card {
        case "ace":
            res = 11
        case "two":
            res = 2
        case "three":
            res = 3
        case "four":
            res = 4
        case "five":
            res = 5
        case "six":
            res = 6
        case "seven":
            res = 7
        case "eight":
            res = 8
        case "nine":
            res = 9
        case "ten":
            res = 10
        case "jack":
            return 10
        case "queen":
            res = 10
        case "king":
            res = 10
        default:
            res = 0
        }

    return res
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var res string
    
    switch {
        case card1 == "ace" && card2 == "ace":
            res = "P"
        case ParseCard(card1) + ParseCard(card2) == 21 && ParseCard(dealerCard) < 10:
            res = "W"
        case ParseCard(card1) + ParseCard(card2) == 21 && ParseCard(dealerCard) >= 10:
            res = "S"
        case ParseCard(card1) + ParseCard(card2) >= 17 && ParseCard(card1) + ParseCard(card2) <= 20:
        	res = "S"
        case ParseCard(card1) + ParseCard(card2) >= 12 && ParseCard(card1) + ParseCard(card2) <= 16 && ParseCard(dealerCard) < 7:
        	res = "S"
        case ParseCard(card1) + ParseCard(card2) >= 12 && ParseCard(card1) + ParseCard(card2) <= 16:
        	res = "H"
        case ParseCard(card1) + ParseCard(card2) < 12:
        	res = "H"
        default:
            res = "H"
    }

    return res
}
