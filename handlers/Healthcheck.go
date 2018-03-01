package handlers

import (
	"encoding/json"
	"net/http"
)

// Healthcheck handles "/healthcheck"
func Healthcheck(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("content-type", "application/json")
	json.NewEncoder(rw).Encode(
		struct {
			Status string `json:"status"`
		}{
			Status: "ok",
		})
}
