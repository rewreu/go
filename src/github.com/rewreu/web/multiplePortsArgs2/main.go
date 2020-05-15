package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func index(w http.ResponseWriter, r *http.Request, mystr string) {
	// fmt.Fprintf(w, "<h1>Port %s is working</h1>", mystr)
	fmt.Fprintf(w, "Successfully connected to port %s\n", mystr)
	dt := time.Now()
	fmt.Fprintf(w, "Server time: %s\n", dt.Format("01-02-2006 15:04:05"))
}

func main() {

	argsWithoutProg := os.Args[1:] // get the port numbers

	for i := 0; i < len(argsWithoutProg); i++ {
		port := argsWithoutProg[i]
		r1 := http.NewServeMux()
		r1.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			index(w, r, port)

		})
		go func() { log.Fatal(http.ListenAndServe(":"+port, r1)) }()
	}

	select {}
}
