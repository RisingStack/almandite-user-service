package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestChain(t *testing.T) {
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

	expectedCalledTimes := 2
	if calledTimes != expectedCalledTimes {
		t.Errorf("handler has not called both middleware: got %d want %d", calledTimes, expectedCalledTimes)
	}
	expectedCalledFirst := "last"
	if calledFirst != expectedCalledFirst {
		t.Errorf("the handler has called the first attached middleware first: got %s want %s", calledFirst, expectedCalledFirst)
	}
}
