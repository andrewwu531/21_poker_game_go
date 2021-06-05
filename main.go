package main

import (
	"bufio"
	"fmt"
	"os"

	"poker21game/card"
	"poker21game/player"
)

func main() {

	var deck = card.GetDecks()
	fmt.Println(deck)
	player.ShowOptions()

	Scanner := bufio.NewScanner(os.Stdin)
	Scanner.Scan()
	var user_option = Scanner.Text()

	for {

		switch user_option {
		case "Hit":
			player.Hit()
		case "Stick":
			fmt.Println("Stick")
		case "Card":
			fmt.Println("Card")
		case "Connect":
			fmt.Println("Connect")
		default:
			fmt.Println("Others")
		}

		if user_option == "Stick" {
			break
		}

		player.ShowOptions()
		Scanner := bufio.NewScanner(os.Stdin)
		Scanner.Scan()
		user_option = Scanner.Text()

	}
}
