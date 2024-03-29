package main

import (
	"fmt"
	"log"
)

type myError struct {
	errType string
	err     error
}

// ==================================================================
func (e *myError) Error() string {
	return fmt.Sprintf("[%s : %s] ", e.errType, e.err.Error())
}

// ==================================================================
//
// ==================================================================
func main() {

	_, err := produceError(true) // Also try with 'false'
	if err != nil {
		log.Println(err)
	} else {
		log.Println("No error!")
	}

}

// ==================================================================
func produceError(b bool) (int, error) {
	if !b {
		return 0, nil
	}

	err := fmt.Errorf("It's % v that I'm just an error value", b)
	return 0, &myError{"critical", err}

}
