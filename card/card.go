package card

type card struct {
	rank int
	suit string
}

func GetDecks() []card {
	// 11 = Jack, 12 = Queen, 13 = King, 14 = Aces
	var decks []card
	for i := 2; i <= 14; i++ {
		for j := 1; j <= 4; j++ {
			switch j {
			case 1:
				decks = append(decks, card{i, "Diamonds"})
			case 2:
				decks = append(decks, card{i, "Clubs"})
			case 3:
				decks = append(decks, card{i, "Hearts"})
			case 4:
				decks = append(decks, card{i, "Spades"})
			}
		}
	}
	return decks
}
