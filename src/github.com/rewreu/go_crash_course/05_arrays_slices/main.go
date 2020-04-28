package main

import "fmt"

func main() {
	// Array
	var fruitArr [2]string

	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// declair and assign
	fruitSlice := []string{"Apple", "Orange", "Grape"}

	fmt.Println(fruitArr)
	fmt.Println(fruitSlice[1:3])
}
