package routes

import (
	"api/src/middlewares"
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
	routes = append(routes, loginRoute)

	for _, route := range routes {
		if route.Authenticate {
			r.HandleFunc(
				route.URI, 
				middlewares.Logger(middlewares.Auth(route.F)),
			).Methods(route.Method)
		}
		r.HandleFunc(route.URI, middlewares.Logger(route.F)).Methods(route.Method)
	}

	return r
}
