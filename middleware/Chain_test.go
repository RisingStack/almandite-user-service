package middleware

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

	var calledTimes int = 0
	var calledFirst string

	createMiddleware := func(name string) func(handler http.HandlerFunc) http.HandlerFunc {
		return func(handler http.HandlerFunc) http.HandlerFunc {
			if calledTimes == 0 {
				calledFirst = name
			}
			calledTimes += 1
			return handler
		}
	}

	firstMiddleware := createMiddleware("first")
	secondMiddleware := createMiddleware("last")

	emptyHandlerFunc := func(w http.ResponseWriter, r *http.Request) {}

	handler := http.HandlerFunc(
		Chain(
			firstMiddleware,
			secondMiddleware,
		)(emptyHandlerFunc),
	)

	handler.ServeHTTP(rr, req)

	if calledTimes != 2 {
		t.Error("handler has not called both middleware")
	}
	if calledFirst != "last" {
		t.Error("the handler has called the first attached middleware first")
	}
}
