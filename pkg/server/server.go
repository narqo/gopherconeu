package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func New() http.Handler {
	r := mux.NewRouter()

	r.Handle("/home", homeHandler()).Methods(http.MethodGet)

	return r
}

func homeHandler() http.Handler {
	h := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, GopherCon EU")
	}
	return http.HandlerFunc(h)
}
