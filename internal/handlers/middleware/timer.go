package middleware

import (
	"log"
	"net/http"
	"time"
)

// Timer middleware logs how much a request has spent in our system
func Timer(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		elapsed := time.Since(start).Nanoseconds() / int64(1e+6)
		log.Printf("Response took: %d ms", elapsed)
	}
}
