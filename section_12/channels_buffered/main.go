package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int, 1) // test with 1 to 3

	for i := 1; i <= 3; i++ {
		go printMessage(c, i)
	}
	time.Sleep(10 * time.Second)

}

// ==============================================================
func printMessage(c chan int, id int) {

	fmt.Printf("ooo %d is waiting for a channel space...\n", id)

	c <- id
	fmt.Printf("=== %d has channel space\n", id)

	time.Sleep(600 * time.Millisecond)
	fmt.Printf("xxx %d has released the channel space\n", id)

	<-c

}
