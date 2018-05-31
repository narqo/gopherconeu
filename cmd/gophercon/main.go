package main

import (
	"fmt"
	"log"
	"net/http"
)

const addr = ":8000"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, GopherCon EU")
	})

	log.Printf("server is running: addr %s\n", addr)

	http.ListenAndServe(addr, nil)
}
