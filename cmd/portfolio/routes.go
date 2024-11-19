package main

import (
	"net/http"

	"github.com/dr4ghs/portfolio/web/middleware"
)

var mux *http.ServeMux

func init() {
	mux = http.NewServeMux()

	registerRouter(mux)
}

func registerRouter(mux *http.ServeMux) {
	mux.HandleFunc("GET /", landingPageHandler)
}

func landingPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)

		return
	}

	middleware.Chain(w, r)
}
