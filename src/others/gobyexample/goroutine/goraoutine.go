package main

import (
    "fmt"
    "time"
)

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}
func l(msg string) {
	fmt.Println(msg)
}

func main() {

    f("direct")

    go f("goroutine")

    go l("going")

    time.Sleep(time.Second)
    fmt.Println("done")
}