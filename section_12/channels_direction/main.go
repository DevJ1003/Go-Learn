package main

import (
	"fmt"
	"math/rand"
	"time"
)

var currentBalance = 200

func main() {

	fmt.Println("Press Enter to stop the program...")
	rand.Seed(time.Now().UnixNano())

	c := make(chan int) // bidirectional

	go credit(c)
	go debit(c)
	go balance(c)

	var input string
	fmt.Scanln(&input) // Press enter to stop the program...

}

// ==============================================================
func credit(c chan int) { // a bi-directional channel
	for i := 0; ; i++ {
		c <- rand.Intn(9) + 1 //random(1-10)
	}
}

// ==============================================================
func debit(c chan<- int) { // a send-only channel
	for i := 0; ; i++ {
		c <- rand.Intn(9) - 10 //random(-10, -1)
	}
}

// ==============================================================
func balance(c <-chan int) { // a receive-only channel
	for {
		num, ok := <-c
		if ok == false {
			fmt.Println("Error:")
			break
		}

		oldBalance := currentBalance
		currentBalance += num //currentBalance = currentBalance + num

		fmt.Printf("=> %d + (%d) = %d\n", oldBalance, num, currentBalance)
		time.Sleep(1 * time.Second)
	}
}
