package handler

import (
	"encoding/json"
	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
)

// A HealthzHandler implements health check endpoint.
type HealthzHandler struct{}

// NewHealthzHandler returns HealthzHandler based http.Handler.
func NewHealthzHandler() *HealthzHandler {
	return &HealthzHandler{}
}

// ServeHTTP implements http.Handler interface.
func (h *HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	message := "OK"
	hr := &model.HealthzResponse{
		Message: message,
	}
	err := json.NewEncoder(w).Encode(hr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
