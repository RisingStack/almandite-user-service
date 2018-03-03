package authentication

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBasicAuthMissing(t *testing.T) {
	req, err := http.NewRequest("GET", "/healthcheck", nil)
	if err != nil {
		t.Fatal(err)
	}

	var isHandlerCalled bool
	rr := httptest.NewRecorder()
	handler := Basic(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

func TestBasicAuthInvalid(t *testing.T) {
	req, err := http.NewRequest("GET", "/healthcheck", nil)
	if err != nil {
		t.Fatal(err)
	}

	var isHandlerCalled bool
	rr := httptest.NewRecorder()
	req.SetBasicAuth("not username", "not password")
	handler := Basic(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	handler := Basic(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

// func TestBasic(t *testing.T) {
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	t.Fatal(string(hashedPassword))
// }
