package authentication

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RisingStack/almandite-user-service/internal/dal/models"
)

type fakeUserRepository struct{}

func (r *fakeUserRepository) GetByID(id int) (*models.User, error) {
	return nil, errors.New("not implemented")
}
func (r *fakeUserRepository) GetByUsername(username string) (*models.User, error) {
	secretUsername := "username"
	secretPassword := "$2a$10$l9FMrAkcGoFv6ghlHq9CSedTt8QO9AeP3KQEaeOR2p3U1eauOSEoO"

	if username == secretUsername {
		return &models.User{
			Username: secretUsername,
			Password: secretPassword,
		}, nil
	}
	return nil, errors.New("Invalid username")
}
func (r *fakeUserRepository) Fetch() (*[]models.User, error) {
	return nil, errors.New("not implemented")
}
func (r *fakeUserRepository) Create(user *models.User) error {
	return errors.New("not implemented")
}
func (r *fakeUserRepository) Update(user *models.User) error {
	return errors.New("not implemented")
}
func (r *fakeUserRepository) Delete(id int) error {
	return errors.New("not implemented")
}

func TestBasicAuthMissing(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/healthcheck", nil)
	if err != nil {
		t.Fatal(err)
	}

	var isHandlerCalled bool
	rr := httptest.NewRecorder()

	authStore := AuthStore{UserRepository: &fakeUserRepository{}}
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

	authStore := AuthStore{UserRepository: &fakeUserRepository{}}
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

	authStore := AuthStore{UserRepository: &fakeUserRepository{}}
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
	authStore := AuthStore{UserRepository: &fakeUserRepository{}}
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
