package authentication

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// Basic authentication middleware
func (store *AuthStore) Basic(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		basicAuthUsername, basicAuthPassword, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "Basic authentication missing", http.StatusUnauthorized)
			return
		}

		user, err := store.UserRepository.GetByUsername(basicAuthUsername)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(basicAuthPassword))
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
