package authentication

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type userFetcher struct{}

func (fetcher *userFetcher) getUserByUsername(username string) (string, error) {
	secretUsername := "username"
	secretPassword := "$2a$10$l9FMrAkcGoFv6ghlHq9CSedTt8QO9AeP3KQEaeOR2p3U1eauOSEoO"

	if username == secretUsername {
		return secretPassword, nil
	}
	return "", errors.New("Invalid username")
}

func TestBasicAuthMissing(t *testing.T) {
	req, err := http.NewRequest("GET", "/healthcheck", nil)
	if err != nil {
		t.Fatal(err)
	}

	var isHandlerCalled bool
	rr := httptest.NewRecorder()

	authStore := AuthStore{userFetcher: &userFetcher{}}
	handler := authStore.Basic(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isHandlerCalled = true
	}))

	handler.ServeHTTP(rr, req)

	if isHandlerCalled {
		t.Error("handler has been called")
	}
	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnauthorized)
	}
	expectedBody := "Basic authentication missing\n"
	if rr.Body.String() != expectedBody {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expectedBody)
	}
}

func TestBasicAuthInvalidUsername(t *testing.T) {
	req, err := http.NewRequest("GET", "/healthcheck", nil)
	if err != nil {
		t.Fatal(err)
	}

	var isHandlerCalled bool
	rr := httptest.NewRecorder()
	req.SetBasicAuth("not username", "not password")

	authStore := AuthStore{userFetcher: &userFetcher{}}
	handler := authStore.Basic(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isHandlerCalled = true
	}))

	handler.ServeHTTP(rr, req)

	if isHandlerCalled {
		t.Error("handler has been called")
	}
	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnauthorized)
	}
	expectedBody := "Invalid credentials\n"
	if rr.Body.String() != expectedBody {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expectedBody)
	}
}

func TestBasicAuthInvalidPassword(t *testing.T) {
	req, err := http.NewRequest("GET", "/healthcheck", nil)
	if err != nil {
		t.Fatal(err)
	}

	var isHandlerCalled bool
	rr := httptest.NewRecorder()
	req.SetBasicAuth("username", "not password")

	authStore := AuthStore{userFetcher: &userFetcher{}}
	handler := authStore.Basic(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isHandlerCalled = true
	}))

	handler.ServeHTTP(rr, req)

	if isHandlerCalled {
		t.Error("handler has been called")
	}
	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnauthorized)
	}
	expectedBody := "Invalid credentials\n"
	if rr.Body.String() != expectedBody {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expectedBody)
	}
}

func TestBasicAuthValid(t *testing.T) {
	req, err := http.NewRequest("GET", "/healthcheck", nil)
	if err != nil {
		t.Fatal(err)
	}

	var isHandlerCalled bool
	rr := httptest.NewRecorder()
	req.SetBasicAuth("username", "password")
	authStore := AuthStore{userFetcher: &userFetcher{}}
	handler := authStore.Basic(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isHandlerCalled = true
	}))

	handler.ServeHTTP(rr, req)

	if !isHandlerCalled {
		t.Error("handler has not been called")
	}
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
