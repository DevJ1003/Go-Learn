package main

import (
	"fmt"
	"sort"
)

func main() {

	n := []int{7, 2, 5, 8, 9, 15, 1}
	fmt.Println(n)

	sort.Sort(sort.IntSlice(n))
	fmt.Println(n)

	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)

	// ==========================================
	s := []string{"Dave", "Rob", "Mike", "Chester", "Brad", "Joe"}
	fmt.Printf("\n%s \n", s)

	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)

}
