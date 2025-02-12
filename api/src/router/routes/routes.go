package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI string
	Method string
	F func(http.ResponseWriter, *http.Request)
	Authenticate bool
}

func Configure(r *mux.Router) *mux.Router {
	routes := userRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.F).Methods(route.Method)
	}

	return r
}
