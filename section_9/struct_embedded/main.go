package main

import "fmt"

type generalInfo struct {
	country    string
	hairColour string
}

type player struct {
	name, sport string
	age         int
	info        generalInfo
}

func main() {

	var player1 player
	player1.name = "Shaquille O` Neal"
	player1.sport = "Basketball"
	player1.age = 50
	player1.info.country = "American"
	player1.info.hairColour = "Black"

	fmt.Println("Player 1: ", player1)

	// ================================================================

	info2 := generalInfo{
		country:    "USA",
		hairColour: "Black",
	}

	player2 := player{
		name:  "Lebron James",
		age:   54,
		sport: "Volleyball",
		info:  info2,
	}

	fmt.Println("Player 2: ", player2)

	// ================================================================

	player3 := player{
		name:  "Serenia Williams",
		sport: "Tennis",
		age:   45,
		info: generalInfo{
			country:    "Africa",
			hairColour: "Blonde",
		},
	}

	fmt.Println("Player 3: ", player3)

	// ================================================================

	player4 := player{"Dhoni", "Cricket", 41, generalInfo{"India", "Black"}}
	fmt.Println("Player 4: ", player4)

	p4 := &player4
	fmt.Println("*p4: ", *p4)
	fmt.Printf("(*p4).name=%s (*p4).info.country=%s\n", (*p4).name, (*p4).info.country)

}
