package main

import (
	"fmt"
	"strings"

	// p1 is alias for "animals" package
	p1 "Projects/Go_Learn/section_13/packages/animals"
	"Projects/Go_Learn/section_13/packages/athletes"
	"Projects/Go_Learn/section_13/packages/shapes1"
	"Projects/Go_Learn/section_13/packages/shapes2"
)

func main() {

	// testCase1()
	// testCase2()
	testCase3()

}

// ==============================================================================
// =====
// ==============================================================================
func testCase1() {

	fmt.Printf("shape1 pkg - Rectangle Area = %d\n", shapes1.Area(shapes1.Length, shapes1.Width))
	fmt.Printf("shape1 pkg - Square Area = %d\n", shapes1.SquareArea(2))

	fmt.Printf("shape2 pkg - Rectangle Area = %d\n", shapes2.Area(2, 3))
	fmt.Printf("shape2 pkg - Square Area = %d\n", shapes2.SquareArea(2))

}

// ==============================================================================
func testCase2() {

	lion := p1.Lion{}
	fmt.Println(lion.Speaks())

	cat := p1.Cat{}
	fmt.Println(cat.Speaks())

	human := p1.Human{}
	fmt.Println(human.Speaks())

}

// ==============================================================================
func testCase3() {

	player1 := athletes.Player{"Anderson Silva", "MMA", 43, athletes.Info{"Brazil", "Black"}}
	fmt.Println("(1) player1:", player1)

	changeAthleteName1(player1)
	fmt.Println("(2) player1:", player1)

	changeAthleteName2(&player1)
	fmt.Println("(3) player1:", player1)

	fmt.Println("(4) player1:", *player1.ToLowerCase())

}

func changeAthleteName1(p athletes.Player) {
	p.Name = "Anderson Silva-Spider"
}

func changeAthleteName2(p *athletes.Player) {
	p.Name = "Anderson Silva-Spider"
	p.Country = strings.ToUpper(p.Country)
}
