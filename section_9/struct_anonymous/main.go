package main

import "fmt"

type player struct {
	string
	int
}

func main() {

	player1 := player{"Shaq", 50}
	fmt.Println("Player 1: ", player1)
	fmt.Printf("player1.int=%d player1.string=%s\n", player1.int, player1.string)

	player2 := player{
		int:    48,
		string: "Lebron",
	}

	fmt.Println("Player2: ", player2)

}
