package routes

import (
	"net/http"

	"github.com/enzocazenave/golang-full-rest-api/internal/handlers"
)

func SetupRoutes(mux *http.ServeMux, handler *handlers.Handler) {
	SetupHealthRoute(mux, handler)
}
