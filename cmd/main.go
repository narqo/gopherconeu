package cmd

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

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

	healthPort := os.Getenv("HEALTH_PORT")
	if healthPort == "" {
		log.Fatal("HEALTH_PORT is empty")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)

	errc := make(chan error, 2)

	heathAddr := net.JoinHostPort("", healthPort)
	hs := server.New(heathAddr, heath.New())

	go func() {
		log.Printf("health checks are running: addr %s\n", heathAddr)
		errc <- hs.Start()
	}()

	addr := net.JoinHostPort("", port)
	s := server.New(addr, routing.New())

	go func() {
		log.Printf("server is running: addr %s, version %s, build %s, buildTime %s\n",
			addr, version.Version, version.Commit, version.BuildTime)
		errc <- s.Start()
	}()

	select {
	case err := <-errc:
		log.Fatal(err)
	case s := <-sigc:
		log.Printf("got %s, stop\n", s)
	}

	s.Stop(ctx)
	hs.Stop(ctx)
}
