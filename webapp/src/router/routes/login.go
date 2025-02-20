package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var routesLogin = []Route {
	{
		URI: "/",
		Method: http.MethodGet,
		F: controllers.LoadLoginScreen,
		Authentication: false,
	},
	{
		URI: "/login",
		Method: http.MethodGet,
		F: controllers.LoadLoginScreen,
		Authentication: false,
	},

}
