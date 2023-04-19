package main

import "fmt"

func main() {

	capitalCity := map[string]string{

		"India":   "Delhi",
		"England": "London",
		"USA":     "Washington",
	}

	fmt.Println(capitalCity["England"])

	// employees := map[string]map[string]string{

	// 	"DJ": map[string]string{
	// 		"Firstname": "Dev",
	// 		"Lastname":  "Joshi",
	// 	},

	// 	"AS": map[string]string{

	// 		"Firstname": "Action",
	// 		"Lastname":  "Sharma",
	// 	},
	// }

	// fmt.Println(employees["DJ"])

}
