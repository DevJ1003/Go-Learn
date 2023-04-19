package main

import "fmt"

func main() {

	var SeasonNo int
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &SeasonNo)

	// SeasonNo := 2

	switch SeasonNo {
	case 1:
		fmt.Println("Spring - ", SeasonNo)

	case 2:
		fmt.Println("Summer - ", SeasonNo)

	case 3:
		fmt.Println("Winter - ", SeasonNo)

	case 4:
		fmt.Println("Monsoon - ", SeasonNo)

	default:
		fmt.Println("a new season - ", SeasonNo)

	}

}
