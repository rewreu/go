package main

import (
	"fmt"
)

func worker(n int, ch chan int) {
	if n < 10 {

		ch <- n

	} else {
		close(ch)
	}

}
func main() {
	ch := make(chan int, 2)

	// ch <- 1
	ch <- 2

	n := 3
	for num := range ch {
		fmt.Println(num)
		worker(n, ch)
		n++
		if n > 12 {
			break
		}
	}

}
