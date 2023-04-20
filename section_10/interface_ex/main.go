package main

import "fmt"

// =================================================================
type rectangle struct {
	w int
	l int
}

func (c *rectangle) area() int {
	return c.w * c.l
}

func (c *rectangle) perim() int {
	return 2 * (c.w * c.l)
}

// =================================================================
type square struct {
	s int
}

func (c *square) area() int {
	return c.s * c.s
}

func (c *square) perim() int {
	return 4 * c.s
}

// =================================================================
type shape interface {
	area() int
	perim() int
}

// =================================================================
func info(s shape) {
	fmt.Printf("area()=%d  perim()=%d", s.area(), s.perim())
}

// =================================================================
func main() {

	r1 := rectangle{2, 3}
	fmt.Printf("r1.area()=%d\n", r1.area())
	fmt.Printf("r1.perim()=%d\n", r1.perim())

	r2 := rectangle{1, 4}
	fmt.Printf("r1.area()=%d\n", r2.area())
	fmt.Printf("r1.perim()=%d\n", r2.perim())

	s1 := square{3}
	fmt.Printf("s1.area()=%d\n", s1.area())
	fmt.Printf("s1.perim()=%d\n", s1.perim())

	s2 := square{7}
	fmt.Printf("s2.area()=%d\n", s2.area())
	fmt.Printf("s2.perim()=%d\n", s2.perim())

	fmt.Printf("\n==========================")
	info(&r1)
	info(&r2)
	info(&s1)
	info(&s2)

}
