package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var homeRoute = Route{
	URI:            "/home",
	Method:         http.MethodGet,
	F:              controllers.LoadHomePage,
	Authentication: true,
}
