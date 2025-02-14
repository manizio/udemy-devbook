package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI: "/usuarios",
		Method: http.MethodPost,
		F: controllers.CreateUser,
		Authenticate: false,
	},

	{
		URI: "/usuarios",
		Method: http.MethodGet,
		F: controllers.SearchAllUsers,
		Authenticate: true,
	},

	{
		URI: "/usuarios/{userID}",
		Method: http.MethodGet,
		F: controllers.SearchUser,
		Authenticate: true,
	},

	{
		URI: "/usuarios/{userID}",
		Method: http.MethodPut,
		F: controllers.UpdateUser,
		Authenticate: true,
	},

	{
		URI: "/usuarios/{userID}",
		Method: http.MethodDelete,
		F: controllers.DeleteUser,
		Authenticate: true,
	},

}
