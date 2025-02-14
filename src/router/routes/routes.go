package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI       string
	Method    string
	Callback  func(http.ResponseWriter, *http.Request)
	IsPrivate bool
}

func SetRoutes(r *mux.Router) *mux.Router {
	routes := notesRoutes

	for _, rota := range routes {
		r.HandleFunc(rota.URI, rota.Callback).Methods(rota.Method)
	}

	return r
}
