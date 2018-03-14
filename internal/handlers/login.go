package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/RisingStack/almandite-user-service/internal/config"
	"github.com/RisingStack/almandite-user-service/internal/dal/models"
	jwt "github.com/dgrijalva/jwt-go"

	"github.com/RisingStack/almandite-user-service/internal/dal"
)

// LoginHandler ...
type LoginHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
}

type loginHandler struct {
	users dal.UserRepository
}

// NewLoginHandler ...
func NewLoginHandler(userRepository dal.UserRepository) LoginHandler {
	return &loginHandler{
		users: userRepository,
	}
}

// Login ...
func (l *loginHandler) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}

	var user models.User

	if err := json.Unmarshal(body, &user); err != nil {
		http.Error(w, "Failed to parse body", http.StatusInternalServerError)
		return
	}

	found, err := l.users.GetByUsername(user.Username)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(found.Password), []byte(user.Password)); err != nil {
		http.Error(w, "Invalid password", http.StatusInternalServerError)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
	})

	signature, err := token.SignedString([]byte(config.GetConfiguration().JwtSigningKey))

	if err != nil {
		http.Error(w, "Failed to sign token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(signature))
}
