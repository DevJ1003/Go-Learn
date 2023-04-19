package main

import "fmt"

func main() {

	serverStatusOk := true

	if serverStatusOk == true {
		fmt.Println("server running")
	} else {
		fmt.Println("not running")
	}

	if s := "NEW"; serverStatusOk == true {
		fmt.Printf("%s server is running good", s)
	} else {
		fmt.Printf("%s server not running", s)
	}

}
