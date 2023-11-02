package main

import (
	"net/http"
)

func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		addHeader(w)
		next.ServeHTTP(w, r)
	})
}

func addHeader(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
}
