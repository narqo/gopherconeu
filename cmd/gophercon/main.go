package main

import (
	"log"
	"net"
	"os"

	"github.com/narqo/gopherconeu/pkg/routing"
	"github.com/narqo/gopherconeu/pkg/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is empty")
	}

	r := routing.New()

	addr := net.JoinHostPort("", port)
	s := server.New(addr, r)

	log.Printf("server is running: addr %s\n", addr)

	log.Fatal(s.Start())
}
