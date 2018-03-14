package authentication

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RisingStack/almandite-user-service/internal/config"
	"github.com/dgrijalva/jwt-go"
)

func TestAuthorizationHeaderMissing(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/secret", nil)
	if err != nil {
		t.Fatal(err)
	}

	var isHandlerCalled bool
	rec := httptest.NewRecorder()

	authStore := AuthStore{}
	handler := authStore.Jwt(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isHandlerCalled = true
	}))

	handler.ServeHTTP(rec, req)

	if isHandlerCalled {
		t.Error("handler has been called")
	}

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("handler expected to return HTTP %v but returned HTTP %v", http.StatusUnauthorized, rec.Code)
	}
}

func TestAuthorizationHeaderInvalid(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/secret", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "thisisaninvalidtoken")

	var isHandlerCalled bool
	rec := httptest.NewRecorder()

	authStore := AuthStore{}
	handler := authStore.Jwt(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isHandlerCalled = true
	}))

	handler.ServeHTTP(rec, req)

	if isHandlerCalled {
		t.Error("handler has been called")
	}

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("handler expected to return HTTP %v but returned HTTP %v", http.StatusUnauthorized, rec.Code)
	}
}

func TestAuthorizationHeaderValid(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/secret", nil)
	if err != nil {
		t.Fatal(err)
	}

	token := jwt.New(jwt.SigningMethodHS256)
	signature, err := token.SignedString([]byte(config.GetConfiguration().JwtSigningKey))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", signature))

	var isHandlerCalled bool
	rec := httptest.NewRecorder()

	authStore := AuthStore{}
	handler := authStore.Jwt(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isHandlerCalled = true
	}))

	handler.ServeHTTP(rec, req)

	if !isHandlerCalled {
		t.Error("handler has not been called")
	}

	if rec.Code != http.StatusOK {
		t.Errorf("handler expected to return HTTP %v but returned HTTP %v", http.StatusOK, rec.Code)
	}
}
