package routes

import (
	"net/http"
	"webapp/src/middlewares"

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
	routes = append(routes, usersRoutes...)
	routes = append(routes, homeRoute)

	for _, route := range routes {
		if route.Authentication {
			router.HandleFunc(
				route.URI, 
				middlewares.Logger(middlewares.Authenticate(route.F)),
			).Methods(route.Method)
		} else {
			router.HandleFunc(
				route.URI,
				middlewares.Logger(route.F),
			).Methods(route.Method)
		}
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", fileServer))
	return router
}
