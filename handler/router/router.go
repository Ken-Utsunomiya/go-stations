package router

import (
	"database/sql"
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler"
	"github.com/TechBowl-japan/go-stations/handler/middleware"
	"github.com/TechBowl-japan/go-stations/service"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()
	hz := handler.NewHealthzHandler()
	svc := service.NewTODOService(todoDB)
	ts := handler.NewTODOHandler(svc)
	meh := handler.NewMockErrorHandler()

	// register routes
	mux.HandleFunc("/healthz", hz.ServeHTTP)
	mux.HandleFunc("/todos", ts.ServeHTTP)
	mux.Handle("/do-panic", middleware.Recovery(meh))

	return mux
}
