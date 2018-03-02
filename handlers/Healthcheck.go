package handlers

import (
	"encoding/json"
	"net/http"
)

// Healthcheck handles "/healthcheck"
func Healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(
		struct {
			Status string `json:"status"`
		}{
			Status: "ok",
		})
}
