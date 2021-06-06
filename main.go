package main

import (
	"bufio"
	"fmt"
	"os"

	"poker21game/card"
	"poker21game/player"
)

func main() {

	var deck = card.GetDeck()
	var hand card.Deck
	score := []int{0}
	previous_games := []int{}
	//1. fmt.Printf("Full Deck %s \n", deck)
	card.Shuffle(deck)
	//1. fmt.Println(deck)
	player.ShowOptions()

	Scanner := bufio.NewScanner(os.Stdin)
	Scanner.Scan()
	var user_option = Scanner.Text()

	for {

		switch user_option {
		case "Hit":
			hand, score, previous_games = card.HitOption(&deck, &hand, &score, &previous_games)
			if score[0] >= 21 {
				hand = nil
				score = []int{0}
				deck = nil
				fmt.Println("New Game -------------------")
				deck = card.GetDeck()
				card.Shuffle(deck)
			}
			// fmt.Printf("\nWhole Deck: %s\n", deck)
			// fmt.Printf("\nCurrent Hand: %s\n", hand)
			// fmt.Printf("\nCurrent Score: %d\n", score[0])
		case "Stick":
			fmt.Println("Stick")
			previous_games = append(previous_games, (score)[0])
			hand = nil
			score = []int{0}
			deck = nil
			fmt.Println("New Game -------------------")
			deck = card.GetDeck()
			card.Shuffle(deck)
		case "Cards":
			fmt.Println("Cards")
			//fmt.Printf("\nWhole Deck: %s\n", deck)
			fmt.Printf("Current Hand: %s\n", hand)
			fmt.Printf("Current Score: %d\n", score[0])
		case "Connect":
			fmt.Println("\nConnect")
			fmt.Printf("\nPrevious Game Scores: %d", previous_games)
			fmt.Printf("\nAverage Scores for all Previous Game: %f\n", average(previous_games))
		case "EndGame":
			os.Exit(3)
		default:
			fmt.Println("Others")
		}
		// player.ShowOptions()
		Scanner := bufio.NewScanner(os.Stdin)
		Scanner.Scan()
		user_option = Scanner.Text()

	}
}

func average(arr []int) float64 {
	total := 0
	for _, v := range arr {
		total += v
	}
	return float64(total) / float64(len(arr))
}
