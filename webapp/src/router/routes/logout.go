package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var logoutRoute = Route{
	URI: "/logout",
	Method: http.MethodGet,
	F: controllers.Logout,
	Authentication: true,
}
