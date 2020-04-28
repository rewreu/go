package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>This port is working</h1>")
}

func main() {

	argsWithoutProg := os.Args[1:] // get the port numbers

	r1 := http.NewServeMux()
	r1.HandleFunc("/", index)

	for i := 0; i < len(argsWithoutProg); i++ {
		port := argsWithoutProg[i]
		go func() { log.Fatal(http.ListenAndServe(":"+port, r1)) }()
	}

	select {}
}
