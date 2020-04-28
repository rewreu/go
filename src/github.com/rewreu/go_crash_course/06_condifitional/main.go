package main

import "fmt"

func main() {
	x := 51
	y := 5
	if x < y {
		fmt.Printf("%d is less than l %d\n", x, y)
	} else if x == y {
		fmt.Printf("%d is equal %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}
}
