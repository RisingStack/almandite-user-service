package middleware

import (
	"log"
	"net/http"
	"os"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "[wapi] ", log.Ldate|log.Ltime|log.LUTC)
}

// Logger middleware is logging every request method and URL
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Printf(`Request arrived: %s "%v"`, r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	}
}
