package main

import (
	"log"
	"net"
	"os"

	"github.com/narqo/gopherconeu/pkg/heath"
	"github.com/narqo/gopherconeu/pkg/routing"
	"github.com/narqo/gopherconeu/pkg/server"
	"github.com/narqo/gopherconeu/version"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is empty")
	}

	healthPort := os.Getenv("HEATH_PORT")
	if healthPort == "" {
		log.Fatal("HEATH_PORT is empty")
	}

	errc := make(chan error, 1)

	go func() {
		heathAddr := net.JoinHostPort("", healthPort)
		hs := server.New(heathAddr, heath.New())

		log.Printf("health checks are running: addr %s\n", heathAddr)

		errc <- hs.Start()
	}()

	go func() {
		addr := net.JoinHostPort("", port)
		s := server.New(addr, routing.New())

		log.Printf("server is running: addr %s, version %s, build %s, buildTime %s\n",
			addr, version.Version, version.Commit, version.BuildTime)

		errc <- s.Start()
	}()

	log.Fatal(<-errc)
}
