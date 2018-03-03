package authentication

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

const secretUsername string = "username"
const secretPassword string = "$2a$10$l9FMrAkcGoFv6ghlHq9CSedTt8QO9AeP3KQEaeOR2p3U1eauOSEoO"

type user struct {
	username string
	password string
}

func Basic(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "Basic authentication missing", http.StatusUnauthorized)
			return
		}

		err := bcrypt.CompareHashAndPassword([]byte(secretPassword), []byte(password))
		if username != secretUsername || err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
