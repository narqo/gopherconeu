package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/narqo/gopherconeu/pkg/server"
)

const addr = ":8000"

func main() {
	r := mux.NewRouter()

	r.Handle("/home", server.HomeHandler()).Methods(http.MethodGet)

	log.Printf("server is running: addr %s\n", addr)

	http.ListenAndServe(addr, r)
}
