package router

import (
	"database/sql"
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler"
	"github.com/TechBowl-japan/go-stations/service"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()
	hz := handler.NewHealthzHandler()
	svc := service.NewTODOService(todoDB)
	ts := handler.NewTODOHandler(svc)

	// register routes
	mux.HandleFunc("/healthz", hz.ServeHTTP)
	mux.HandleFunc("/todos", ts.ServeHTTP)

	return mux
}
