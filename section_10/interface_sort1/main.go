package main

import (
	"fmt"
	"sort"
)

func main() {

	n := []int{7, 2, 5, 8, 9, 15, 1}
	fmt.Println(n)
	sort.Ints(n)
	fmt.Println(n)

	// ==================================================
	s := []string{"Dave", "Rob", "Mike", "Chester", "Brad", "Joe"}
	fmt.Println(s)
	sort.Strings(s)
	fmt.Println(s)

}
