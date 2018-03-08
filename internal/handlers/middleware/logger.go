package middleware

import (
	"log"
	"net/http"
)

// Logger middleware is logging every request method and URL
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf(`Request arrived: %s "%v"`, r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	}
}
