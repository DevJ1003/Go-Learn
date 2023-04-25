package main

import (
	"fmt"
)

func main() {

	// c := make(chan int) //declaring & defining using short hand declaration

	var c chan int
	// if c == nil {
	// 	c = make(chan int)
	// }

	fmt.Printf("Type of c: %T", c)
	close(c)

}
