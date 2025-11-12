package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func respondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, `{"error":"failed to encode response"}`, http.StatusInternalServerError)
	}
}

func respondErr(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{
		"code":  strconv.Itoa(status),
		"error": message,
	})
}
