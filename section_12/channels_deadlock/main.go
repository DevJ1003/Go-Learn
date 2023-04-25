package main

func main() {

	c := make(chan string)
	c <- "No one likes me!" //fatal error: all goroutines are asleep - deadlock!

}
