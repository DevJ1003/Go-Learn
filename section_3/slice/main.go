package main

import (
	"fmt"
)

func main() {

	s := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	fmt.Println(s)

	slc1 := s[1:3]
	fmt.Println(slc1)

	slc2 := s[5:]
	fmt.Println(slc2)

	join := append(slc1, slc2...)
	fmt.Println(join)

	fmt.Println(len(join))

}
