package rest

import (
	"encoding/json"
	"net/http"
)

// HealthCheck is a liveness check: if the process can respond at all, it's up
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
