package heath

import (
	"fmt"
	"net/http"
)

func New() http.Handler {
	h := http.NewServeMux()
	h.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Ok")
	})
	h.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Ok")
	})
	return h
}
