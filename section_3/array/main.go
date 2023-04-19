package main

import "fmt"

func main() {

	// case 1
	// temp := [...]int{3, 2, -5, 4}
	// fmt.Println(temp)

	// case 2
	// var nums [3]int = [3]int{1, 2, 3}
	// fmt.Printf("nums=%v type=%T len=%d\n", nums, nums, len(nums))

	// case 3
	x := [3]float32{2.1, 3.2, 4.7}

	fmt.Println(x)

	var total float32

	for _, val := range x {

		total += val

	}

	fmt.Println(total)

}
