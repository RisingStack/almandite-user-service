package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthcheck(t *testing.T) {
	req, err := http.NewRequest("GET", "/healthcheck", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Healthcheck)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedContentType := "application/json"
	if rr.HeaderMap.Get("content-type") != expectedContentType {
		t.Errorf("handler returned unexpected content type: got %v want %v",
			rr.Body.String(), expectedContentType)
	}

	expectedBody := "{\"status\":\"ok\"}\n"
	if rr.Body.String() != expectedBody {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expectedBody)
	}
}
