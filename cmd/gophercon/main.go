package main

import (
	"log"
	"net/http"

	"github.com/narqo/gopherconeu/pkg/server"
)

const addr = ":8000"

func main() {
	r := server.New()

	log.Printf("server is running: addr %s\n", addr)

	http.ListenAndServe(addr, r)
}
