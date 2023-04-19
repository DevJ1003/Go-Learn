package main

import "fmt"

func main() {

	// x := 2

	// for i := 1; i <= 10; i++ {

	// 	fmt.Printf("%2d * %d = %d\n", i, x, i*x)

	// }

	for i := 0; i <= 10; i++ {

		if i == 2 {
			continue
		}

		if i == 8 {
			break
		}

		fmt.Println(i)

	}

}
