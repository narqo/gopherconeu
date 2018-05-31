package server

import (
	"fmt"
	"net/http"
)

func HomeHandler() http.Handler {
	h := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, GopherCon EU")
	}
	return http.HandlerFunc(h)
}
