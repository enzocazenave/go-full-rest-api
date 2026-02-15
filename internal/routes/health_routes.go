package routes

import (
	"net/http"

	"github.com/enzocazenave/golang-full-rest-api/internal/handlers"
)

func SetupHealthRoute(mux *http.ServeMux, handler *handlers.Handler) {
	mux.HandleFunc("/health", handler.HealthHandler())
}
