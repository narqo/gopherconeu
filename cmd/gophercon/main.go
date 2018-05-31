package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const addr = ":8000"

func main() {
	r := mux.NewRouter()

	r.Handle("/home", homeHandler()).Methods(http.MethodGet)

	log.Printf("server is running: addr %s\n", addr)

	http.ListenAndServe(addr, r)
}

func homeHandler() http.Handler {
	h := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, GopherCon EU")
	}
	return http.HandlerFunc(h)
}
