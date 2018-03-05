package authentication

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (store *AuthStore) Basic(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		basicAuthUsername, basicAuthPassword, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "Basic authentication missing", http.StatusUnauthorized)
			return
		}

		hashedPassword, err := store.userFetcher.getUserByUsername(basicAuthUsername)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(basicAuthPassword))
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
