package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI            string
	Method         string
	F              func(http.ResponseWriter, *http.Request)
	Authentication bool
}

func Configure(router *mux.Router) *mux.Router {
	routes := routesLogin

	for _, route := range routes {
		router.HandleFunc(route.URI, route.F).Methods(route.Method)
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", fileServer))
	return router
}
