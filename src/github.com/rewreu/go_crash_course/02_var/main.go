package main

import "fmt"

func main() {
	// name := "John"
	// email := "John@gmail.com"
	name, email := "John", "John@gmail.com"
	var age int32 = 37
	const isCool = true
	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", isCool) // print the type of isCool
}
