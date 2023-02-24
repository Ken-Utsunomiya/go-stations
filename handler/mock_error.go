package handler

import (
	"net/http"
)

// A HealthzHandler implements health check endpoint.
type MockErrorHandler struct{}

// NewHealthzHandler returns HealthzHandler based http.Handler.
func NewMockErrorHandler() *MockErrorHandler {
	return &MockErrorHandler{}
}

// ServeHTTP implements http.Handler interface.
func (h *MockErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	panic("Panic triggered!!")
}
