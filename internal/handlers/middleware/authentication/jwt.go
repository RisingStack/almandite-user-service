package authentication

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/RisingStack/almandite-user-service/internal/config"
	"github.com/dgrijalva/jwt-go"
)

// Jwt authentication middleware
func (store *AuthStore) Jwt(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")

		if auth == "" {
			http.Error(w, "Missing Authorization token", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(auth, "Bearer ") {
			http.Error(w, "Invalid Authorization schema", http.StatusUnauthorized)
			return
		}

		tokenString := auth[7:len(auth)]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method")
			}
			return []byte(config.GetConfiguration().JwtSigningKey), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid Authorization header: "+err.Error(), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
