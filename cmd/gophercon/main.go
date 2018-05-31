package main

import (
	"log"
	"net/http"
	"os"

	"github.com/narqo/gopherconeu/pkg/server"
)

func main() {
	addr := os.Getenv("ADDR")
	if addr == "" {
		log.Fatal("ADDR wasn't set")
	}

	r := server.New()

	log.Printf("server is running: addr %s\n", addr)

	log.Fatal(http.ListenAndServe(addr, r))
}
