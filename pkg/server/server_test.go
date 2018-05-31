package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNew(t *testing.T) {
	handler := New()

	testsrv := httptest.NewServer(handler)
	defer testsrv.Close()

	cases := []struct {
		path       string
		wantStatus int
	}{
		{
			"/home",
			http.StatusOK,
		},
		{
			"/blah",
			http.StatusNotFound,
		},
	}

	for _, tc := range cases {
		t.Run(tc.path, func(t *testing.T) {
			res, err := http.Get(testsrv.URL + tc.path)
			if err != nil {
				t.Fatal(err)
			}
			if got, want := res.StatusCode, tc.wantStatus; got != want {
				t.Errorf("GET status %v, want %v", got, want)
			}
		})
	}
}
