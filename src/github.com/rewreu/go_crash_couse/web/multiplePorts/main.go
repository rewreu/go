package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About</h1>")
}

func main() {

	r2 := http.NewServeMux()
	r2.HandleFunc("/", index)

	go func() { log.Fatal(http.ListenAndServe(":8001", r2)) }()
	go func() { log.Fatal(http.ListenAndServe(":8002", r2)) }()
	select {}
}
