package card

import (
	"fmt"
	"math/rand"
)

type Card struct {
	Type string
	Suit string
}

type Deck []Card

func GetDeck() (deck Deck) {

	types := []string{"Two", "Three", "Four", "Five", "Six", "Seven",
		"Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}

	suits := []string{"Heart", "Diamond", "Club", "Spade"}

	for i := 0; i < len(types); i++ {
		for n := 0; n < len(suits); n++ {
			card := Card{
				Type: types[i],
				Suit: suits[n],
			}
			deck = append(deck, card)
		}
	}
	return deck
}

func Shuffle(d Deck) Deck {

	for i := 1; i < len(d); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			d[r], d[i] = d[i], d[r]
		}
	}
	return d
}

func HitOption(d *Deck, h *Deck, score *[]int, previous_games *[]int) (Deck, []int, []int) {

	first_element := (*d)[0]
	card := Card{
		Type: first_element.Type,
		Suit: first_element.Suit,
	}
	*h = append(*h, card)
	*d = (*d)[1:]
	// fmt.Printf("Card Dealt: %s %s \n", first_element.Type, first_element.Suit)
	// fmt.Printf("Current Hand: %s", *h)

	switch first_element.Type {
	case "Two":
		*score = append(*score, 2)
	case "Three":
		*score = append(*score, 3)
	case "Four":
		*score = append(*score, 4)
	case "Five":
		*score = append(*score, 5)
	case "Six":
		*score = append(*score, 6)
	case "Seven":
		*score = append(*score, 7)
	case "Eight":
		*score = append(*score, 8)
	case "Nine":
		*score = append(*score, 9)
	case "Ten", "Jack", "Queen", "King":
		*score = append(*score, 10)
	case "Ace":
		if (*score)[0]+11 > 21 {
			*score = append(*score, 1)
		} else {
			*score = append(*score, 11)
		}
		// fmt.Printf("Ace value = %d", *score)
		// fmt.Printf("Ace value2 = %d", (*score)[0])
	}

	sum := sum_scores(&*score)
	// fmt.Printf("score list = %d", sum)
	(*score)[0] = sum
	fmt.Printf("score list = %d\n", *score)
	fmt.Printf("score list first = %d\n", (*score)[0])

	if (*score)[0] == 21 {
		fmt.Println("Congragulation, you hit 21")
		*previous_games = append(*previous_games, (*score)[0])
	}
	if (*score)[0] > 21 {
		fmt.Println("Unfortunately, you hit above 21. You have lost this round of the game")
		*previous_games = append(*previous_games, (*score)[0])
	}

	return *h, *score, *previous_games
}

func sum_scores(array *[]int) int {
	result := 0
	for i := 1; i < len(*array); i++ {
		result += (*array)[i]
	}
	return result
}
