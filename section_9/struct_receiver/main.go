package main

import "fmt"

type movie struct {
	name  string
	movie string
}

func (m movie) fullInfo() string {
	return m.name + " => " + m.movie
}

func main() {

	m1 := movie{"Fast And Furious", "Vin Diesel"}
	m2 := movie{"Avengers", "Robert Downey Jr."}

	fmt.Println(m1.fullInfo())
	fmt.Println(m2.fullInfo())

}
